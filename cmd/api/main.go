package main

import (

	// "github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/routes"
	"fmt"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/db"
)

func init() {
	config.LoadConfig()
}

func main() {
	// TODO: Place behind arguments flag.
	// routes.Run(API_TOKEN)
	db.PopulateDB()

	fmt.Println(config.DatabaseConfig.StagingURI)
}
