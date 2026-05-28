package env

import (
	"errors"
	"fmt"
	"log"
	"os"

	u "github.com/joho/godotenv"
)

var (
	envERR = errors.New("error loading DB details")
)
func Config()(string,error ){
	err := u.Load()
	if err != nil{
		log.Fatal(envERR)
	}
	dbCreds := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dbCreds, nil
}