package db

import (
	"fmt"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/utility"
)

func PopulateDB() {
	teams := utility.FetchTeamData()

	fmt.Println("Data Fetched")
	for _, team := range teams.Data {
		fmt.Println(team.Name)
	}
}
