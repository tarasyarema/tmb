package main

import (
	"os"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
)

var base_url string
var api_creds Creds

func loadEnv() {
	base_url = "https://api.tmb.cat/v1"
	api_creds.AppId = os.Getenv("TMB_ID")
	api_creds.AppKey = os.Getenv("TMB_KEY")

	fmt.Println("Environment loaded!")
}