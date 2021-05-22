package telegram

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/Yash-Garg/git-release-tracker/constants"
	"github.com/Yash-Garg/git-release-tracker/utils"
	"log"
)

func Notify(b *gotgbot.Bot, ctx *ext.Context) error {
	for i := range constants.RepoList {
		url := fmt.Sprintf(`https://api.github.com/repos/%s/releases/latest`, constants.RepoList[i])
		data := utils.GetJsonData(url)
		// Just for testing purposes [WIP]
		_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Release ID - %d", data.ID), &gotgbot.SendMessageOpts{
			ParseMode: "html",
		})
		if err != nil {
			log.Println("Failed to send: " + err.Error())
		}
	}
	return nil
}
