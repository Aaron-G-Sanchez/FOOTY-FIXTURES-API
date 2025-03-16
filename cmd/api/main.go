package main

import (

	// "github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/routes"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/db"
)

func init() {
	config.LoadConfig()
}

func main() {
	// TODO: Pass desired enviroment into db.Connect via command line flags.
	database := db.Connect(config.DatabaseConfig.StagingURI)
	defer database.Close()

	// TODO: Place behind command line flags.
	db.PopulateDB(database)

	// routes.Run(API_TOKEN)
}
