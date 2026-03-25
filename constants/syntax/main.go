package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit //MaxCacheSize is the maximum number of items that can be stored in the cache, which is 10 times the GlobalLimit
const (                                   //CacheKeyBook is the prefix for book keys in the cache, CacheKeyCD is the prefix for cd keys in the cache
	CacheKeyBook = "book"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string { //key is the key to get the value from the cache
	if val, ok := cache[key]; ok { //check if the key exists in the cache
		return val
	}
	return cache[key]
}

func cacheSet(key, val string) { //key is the key to set the value, val is the value to be set
	if len(cache)+1 >= MaxCacheSize { //check if the cache size exceeds the maximum limit
		return
	}
	cache[key] = val
}
func GetBook(isbn string) string { //isbn is the key to get the book name
	return cacheGet(CacheKeyBook + isbn) //CacheKeyBook is the prefix to identify that the key is for book
}
func SetBook(isbn, name string) { //isbn is the key to set the book name
	cacheSet(CacheKeyBook+isbn, name) //CacheKeyBook is the prefix to identify that the key is for book
}
func GetCD(sku string) string { //sku is the key to get the cd title
	return cacheGet(CacheKeyCD + sku) //CacheKeyCD is the prefix to identify that the key is for cd
}
func SetCD(sku, title string) { //sku is the key to set the cd title
	cacheSet(CacheKeyCD+sku, title) //CacheKeyCD is the prefix to identify that the key is for cd
}
func main() {
	cache = make(map[string]string)                  //initialize the cache map
	SetBook("1234-5678", "Get Ready To Go")          //isbn is the key to set the book name, "Get Ready To Go" is the book name
	SetCD("1234-5678", "Get Ready To Go Audio Book") //sku is the key to set the cd title, "Get Ready To Go Audio Book" is the cd title
	fmt.Println("Book:", GetBook("1234-5678"))
	fmt.Println("CD :", GetCD("1234-5678"))
}
