package spotify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Test() {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/me/top/tracks", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ""))
	queryString := request.URL.Query()
	queryString.Add("time_range", "long_term")

	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	_ = fmt.Sprint(result)
}