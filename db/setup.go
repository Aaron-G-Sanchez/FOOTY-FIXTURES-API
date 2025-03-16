package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aaron-g-sanchez/PROJECTS/FOOTY-FIXTURES-DISCORD-BOT-API/utility"
)

func PopulateDB(database *sql.DB) {
	teams := utility.FetchTeamData()

	err := database.Ping()
	if err != nil {
		log.Fatal("DB not responding.")
	}

	// TODO: Write team data to DB.
	fmt.Println(teams)
}
