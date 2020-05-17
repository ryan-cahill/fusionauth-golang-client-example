# FusionAuth Golang Client Example

This is an example of using [FusionAuth's Golang client library](https://github.com/FusionAuth/go-client).

## Setup Intructions

1. Go through the [FusionAuth 5 minute setup](https://fusionauth.io/docs/v1/tech/5-minute-setup-guide) to get your FusionAuth server running locally.
1. Create a `.env` file at the root of the directory and fill with the following values:
   1. `FA_CLIENT_ID`
   1. `FA_CLIENT_SECRET`
   1. `FA_API_KEY`
1. Create a "Golang client example" application, copy the 'Client ID' and 'Client Secret' values to your `.env` file.
1. Create an FusionAuth API Key and copy the value to your `.env` file.
1. Run `go run .` and navigate to `localhost:8080` in your browser to test out the application.
