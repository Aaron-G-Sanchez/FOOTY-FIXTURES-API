package main

import (

	// "github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/routes"

	"fmt"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/db"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/utility"
)

func init() {
	config.LoadConfig()
}

func main() {
	run()
}

func run() {
	// TODO: Pass desired enviroment into db.Connect via command line flags.
	// Default enviroment will be staging.
	database := db.Connect(config.DatabaseConfig.StagingURI)
	defer database.Close()

	// routes.Run(API_TOKEN)

	data := utility.FetchTeamData()

	teams := data.Data

	for _, team := range teams {
		fmt.Println(team.Id, team.Name)
	}
}
