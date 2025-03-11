package utility

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var urls = map[string]string{
	"teamsBySeasonId": "https://api.sportmonks.com/v3/football/teams/seasons/24962?api_token=%v",
}

func FetchTeamData(token string) {
	client := &http.Client{}

	endpoint := fmt.Sprintf(urls["teamsBySeasonId"], token)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("err: ", err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error fetching teams: ", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading body: ", err)
	}

	fmt.Println(string(body))
}
