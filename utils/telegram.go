package utils

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/Yash-Garg/git-release-tracker/commands"
	"github.com/Yash-Garg/git-release-tracker/constants"
	"net/http"
)

func RunBot() {
	b, err := gotgbot.NewBot(constants.BotToken, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("Failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("start", commands.Start))

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("Failed to start polling: " + err.Error())
	}

	fmt.Printf("%s has been started...\n", b.User.Username)
	updater.Idle()
}
