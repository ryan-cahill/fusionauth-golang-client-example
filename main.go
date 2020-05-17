package main

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ClientID string
	ClientSecret string
	ApiKey string
)

func main() {
	godotenv.Load(".env")
	ClientID = os.Getenv("FA_CLIENT_ID")
	ClientSecret = os.Getenv("FA_CLIENT_SECRET")
	ApiKey = os.Getenv("FA_API_KEY")

	r := setupRouter()
	// Listen and Serve on 0.0.0.0:8080
	r.Run(":8080")
}
