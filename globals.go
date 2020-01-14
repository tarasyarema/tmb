package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var baseURL string
var apiCreds Creds

var port string

func loadEnv() {
	baseURL = "https://api.tmb.cat/v1"

  apiCreds.AppID = os.Getenv("TMB_ID")
	apiCreds.AppKey = os.Getenv("TMB_KEY")

  port = os.Getenv("PORT")

	fmt.Println("Environment loaded!")
}
