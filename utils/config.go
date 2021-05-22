package utils

import (
	"github.com/Yash-Garg/git-release-tracker/constants"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No file named .env found")
	}
}

func GetEnvVars() {
	constants.ChatID = os.Getenv("CHAT_ID")
	if len(constants.ChatID) == 0 {
		log.Println("CHAT_ID is not set")
	}

	constants.BotToken = os.Getenv("BOT_TOKEN")
	if len(constants.BotToken) == 0 {
		log.Println("BOT_TOKEN is not set")
	}

	repoList := os.Getenv("REPO_LIST")
	if len(repoList) == 0 {
		log.Println("REPO_LIST is not set")
	}

	repos := strings.Split(repoList, ",")
	for _, s := range repos {
		constants.RepoList = append(constants.RepoList, strings.TrimSpace(s))
	}
}
