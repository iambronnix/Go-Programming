package main

import (
	"fmt"
	"os"
)

type locale struct {
	lang    string
	country string
}

func getLocale() []locale {
	var en_US locale
	var en_CN locale
	var fr_CN locale
	var fr_FR locale
	var ru_RU locale

	en_US.lang = "en_US"
	en_US.country = "United States"
	en_CN.lang = "en_CN"
	en_CN.country = "Canada"
	fr_CN.lang = "fr_CN"
	fr_CN.country = "Canada"
	fr_FR.lang = "fr_FR"
	fr_FR.country = "France"
	ru_RU.lang = "ru_RU"
	ru_RU.country = "Russia"

	return []locale{en_US, en_CN, fr_CN, fr_FR, ru_RU}
}
func main() {
	validate := getLocale()
	if len(os.Args) < 2 {
		fmt.Println("No arguments passed into the terminal")
		os.Exit(1)
	}
	passedArg := struct {
		lang    string
		country string
	}{
		os.Args[1],
		os.Args[2],
	}
	for i := 0; i < len(validate); i++ {
		if passedArg == validate[i] {
			fmt.Printf("Locale passed is supported:%+v\n", validate[i])
			os.Exit(0)
		}
		fmt.Printf("Locale not supported:%+v\n", passedArg)
		os.Exit(1)
	}
}
