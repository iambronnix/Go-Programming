package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)
const IssueURL = "https://api.github.com/search/issues"
type Issue struct{
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User string
	CreatedAt time.Time `json:"created_at"`
	Body string
	
}
type IssuesSearchResult struct{
	TotalCount int `json:"total_count"`
	Items []*Issue
}
type User struct{
	Login string
	HTMLURL string `json:"html_url"`
}
func main(){
	result, err := SearchIssues(os.Args[1:])
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.Totalcount)
	for _, item := range result.Items{
		fmt.Printf("#%-5d	%9.9s %.55s\n",
   item.Number, item.User.Login,item.Title )
	}
}
// SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string)(*IssuesSearchResult,error){
	q := url.QueryEscape(strings.Join(terms," "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil{
		return nil, err
	}
	//must close resp.Body on all execution paths
	if resp.StatusCode != http.StatusOK{
		 resp.Body.Close()
			return nil, fmt.Errorf("search query failed: %s", resp.Status)
		}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result);err != nil{//uses streaming decoder 
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
	
}