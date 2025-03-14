package main

import (
	"fmt"

	// "github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/routes"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/config"
	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/utility"
)

func init() {
	config.LoadConfig()
}

func main() {
	// routes.Run(API_TOKEN)
	// TODO: Remove after testing.
	teams := utility.FetchTeamData()

	for _, team := range teams.Data {
		fmt.Println(team)
	}
}
