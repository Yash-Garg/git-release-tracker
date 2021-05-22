package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Yash-Garg/git-release-tracker/models"
)

func GetJsonData(ApiUrl string) models.Release {
	response, _ := http.Get(ApiUrl)
	if response.StatusCode != 200 {
		log.Println("No releases found")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("ERROR: ", err)
		}
	}(response.Body)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	var release models.Release
	err = json.Unmarshal(data, &release)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	return release
}
