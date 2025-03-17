package main

import (

	// "github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/routes"

	"log"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/db"
)

func init() {
	config.LoadConfig()
}

func main() {
	run()
}

func run() {
	// TODO: Pass desired enviroment into db.Connect via command line flags.
	database := db.Connect(config.DatabaseConfig.StagingURI)
	defer database.Close()

	// TODO: Place behind command line flags.
	err := db.PopulateDB(database)
	if err != nil {
		log.Fatal("Error populating database: ", err)
	}
	// routes.Run(API_TOKEN)
}
