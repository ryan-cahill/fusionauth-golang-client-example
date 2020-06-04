package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	ClientID     string
	ClientSecret string
	ApiKey       string
	FAHost       string
	FAPort       string
)

func main() {
	godotenv.Load(".env")
	ClientID = os.Getenv("FA_CLIENT_ID")
	ClientSecret = os.Getenv("FA_CLIENT_SECRET")
	ApiKey = os.Getenv("FA_API_KEY")
	FAHost = os.Getenv("FA_HOST")
	FAPort = os.Getenv("FA_PORT")

	publicPort := os.Getenv("PUBLIC_PORT")

	r := setupRouter()
	// Listen and Serve on 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%s", publicPort))
}
