package main

import (
	"github.com/Yash-Garg/git-release-tracker/telegram"
	"github.com/Yash-Garg/git-release-tracker/utils"
)

func main() {
	utils.GetEnvVars()
	telegram.RunBot()
}
