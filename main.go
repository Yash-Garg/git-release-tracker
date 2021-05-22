package main

import (
	"fmt"
	"log"
	"net/http"
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
	b, err := gotgbot.NewBot(constants.BotToken, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("Failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: false})
	if err != nil {
		panic("Failed to start polling: " + err.Error())
	}

	fmt.Printf("%s has been started...", b.User.Username)

	for {
		for i := range constants.RepoList {
			url := fmt.Sprintf(`https://api.github.com/repos/%s/releases/latest`, constants.RepoList[i])
			data := utils.GetJsonData(url)
			message := fmt.Sprintf(`Requested ID: %d`, data.ID)
			_, err := b.SendMessage(constants.ChatID, message, &gotgbot.SendMessageOpts{ParseMode: "MarkdownV2"})
			if err != nil {
				log.Println("ERROR: ", err)
			}
		}
		time.Sleep(time.Second * 30)
	}
}
