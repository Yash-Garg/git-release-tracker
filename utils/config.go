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
		log.Fatalln("No file named .env found")
		os.Exit(1)
	}
}

func GetEnvVars() {
	chatId, _ := strconv.ParseInt(os.Getenv("CHAT_ID"), 0, 64)
	constants.ChatID = chatId
	if chatId == 0 {
		log.Fatalln("CHAT_ID is not set")
		os.Exit(1)
	}

	constants.BotToken = os.Getenv("BOT_TOKEN")
	if len(constants.BotToken) == 0 {
		log.Fatalln("BOT_TOKEN is not set")
		os.Exit(1)
	}

	repoList := os.Getenv("REPO_LIST")
	if len(repoList) == 0 {
		log.Fatalln("REPO_LIST is not set")
		os.Exit(1)
	}

	repos := strings.Split(repoList, ",")
	for _, s := range repos {
		constants.RepoList = append(constants.RepoList, strings.TrimSpace(s))
	}
}
