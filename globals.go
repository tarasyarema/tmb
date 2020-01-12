package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var baseURL string
var apiCreds Creds

func loadEnv() {
	baseURL = "https://api.tmb.cat/v1"
	apiCreds.AppID = os.Getenv("TMB_ID")
	apiCreds.AppKey = os.Getenv("TMB_KEY")

	fmt.Println("Environment loaded!")
}
