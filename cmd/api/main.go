package main

import (
	"log"
	"os"

	// "github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/routes"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/utility"
	"github.com/joho/godotenv"
)

var API_TOKEN string

func init() {
	// TODO: Move to config directory.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Set the API token.
	API_TOKEN = os.Getenv("API_TOKEN")
}

func main() {
	// routes.Run(API_TOKEN)
	// TODO: Remove after testing.
	utility.FetchTeamData(API_TOKEN)
}
