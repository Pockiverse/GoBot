package main

import (
	"fmt"

	"github.com/evanstan/GoBot/bot"
	"github.com/evanstan/GoBot/config"
)

func main() {
	err := Config.LoadConfigFile()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := &GoBot.GoBot{
		Token:    Config.Token,
		Prefix:   Config.Prefix,
		EmojiDir: Config.EmojiDir,
	}
	
	s.Start()

	<-make(chan struct{})
	return
}