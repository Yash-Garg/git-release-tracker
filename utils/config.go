package utils

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Yash-Garg/git-release-tracker/constants"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No file named .env found")
	}
}

func GetEnvVars() {
	chatId, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 0, 64)
	constants.ChatID = chatId
	if chatId == 0 {
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
