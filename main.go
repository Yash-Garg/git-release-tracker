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

			fileName := strings.Replace(repo, "/", "_", 1)
			lastUpdateID := utils.GetLastID(fileName)
			repoURL := "https://github.com/" + repo

			changelog := data.Body
			if changelog == "" {
				changelog = "No changes specified"
			}

			message := fmt.Sprintf("<b>New <a href='%s'>%s</a> release detected !</b>\n\n<b>Author : </b><a href='%s'>%s</a>\n<b>Release Name : </b><code>%s</code>\n<b>Release Tag : </b><a href='%s'>%s</a>\n<b>Branch : </b><code>%s</code>\n<b>Changelog : </b><code>%s</code>", repoURL, strings.Split(repo, "/")[1], data.Author.HTMLURL, data.Author.Login, data.Name, data.HTMLURL, data.TagName, data.TargetCommitish, changelog)

			if lastUpdateID != data.ID {
				_, err := b.SendMessage(constants.ChatID, message, &gotgbot.SendMessageOpts{ParseMode: "HTML", DisableWebPagePreview: true})
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
		time.Sleep(time.Minute * 1)
	}
}
