package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func CreateFile(fileName string, data string) {
	path := fmt.Sprintf(`temp/%s.txt`, fileName)
	err := ioutil.WriteFile(path, []byte(data), 0755)
	if err != nil {
		log.Println("Unable to write file: ", err)
	}
}

func GetLastID(fileName string) int64 {
	path := fmt.Sprintf(`temp/%s.txt`, fileName)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error reading file: ", err)
	}

	releaseID, _ := strconv.Atoi(string(data))
	return int64(releaseID)
}
