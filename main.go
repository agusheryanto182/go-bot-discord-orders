package main

import (
	"fmt"

	"github.com/agusheryanto182/go-bot-discord-orders/bot"
	"github.com/agusheryanto182/go-bot-discord-orders/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
