package main

import (
	"fmt"
	"os"
)

type locale struct {
	language string
	country  string
}

func getLocale() []locale {
	var en_US locale
	var en_CN locale
	var fr_CN locale
	var fr_FR locale
	var ru_RU locale
	en_US.language = "en_US"
	en_US.country = "United States"
	en_CN.language = "en_CN"
	en_CN.country = "Canada"
	fr_CN.language = "fr_CN"
	fr_CN.country = "France"
	fr_FR.language = "fr_FR"
	fr_FR.country = "France"
	ru_RU.language = "ru_RU"
	ru_RU.country = "Russia"
	return []locale{en_US, en_CN, fr_CN, fr_FR, ru_RU}
}
func main() {
	parsedArg := struct { //anonymous struct
		language string
		country  string
	}{os.Args[1],
		os.Args[2]}

	validate := getLocale()
	if len(os.Args) < 2 {
		fmt.Println("No language parsed")
		os.Exit(1)
	}
	for i := 0; i < len(validate); i++ {

		if parsedArg == validate[i] {
			fmt.Printf("%s is supported:%s\n", os.Args[1], validate[i])
			os.Exit(0)
		}
	}
	fmt.Printf("locale not supported:\t%s\n", os.Args[1])
	os.Exit(0)
}
