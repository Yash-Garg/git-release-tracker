package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/Yash-Garg/git-release-tracker/constants"
	"github.com/Yash-Garg/git-release-tracker/utils"
)

func init() {
	utils.GetEnvVars()
}

func main() {
	bot, err := gotgbot.NewBot(constants.BotToken, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		log.Panicln("Failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)

	err = updater.StartPolling(bot, &ext.PollingOpts{DropPendingUpdates: false})
	if err != nil {
		log.Panicln("Failed to start polling: " + err.Error())
	}

	log.Printf("%s has been started...\n", bot.User.FirstName)
	regularCheck(bot)

}

func regularCheck(b *gotgbot.Bot) {
	for {
		for i := range constants.RepoList {
			repo := constants.RepoList[i]
			url := fmt.Sprintf(`https://api.github.com/repos/%s/releases/latest`, repo)
			data := utils.GetJsonData(url)
			message := fmt.Sprintf(`Requested ID: %d`, data.ID)
			fileName := strings.Replace(repo, "/", "_", 1)
			lastUpdateID := utils.GetLastID(fileName)

			if lastUpdateID != data.ID {
				_, err := b.SendMessage(constants.ChatID, message, &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
				if err != nil {
					log.Fatalln("ERROR: ", err)
				} else {
					log.Printf("Release Sent (%s) - %d", repo, data.ID)
					utils.CreateFile(fileName, strconv.Itoa(int(data.ID)))
				}
			} else {
				log.Printf(`%s is up to date!`, repo)
			}
		}
		time.Sleep(time.Second * 30)
	}
}
