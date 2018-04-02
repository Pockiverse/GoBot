package main

import (
	"fmt"
	"./config"
	"./bot"
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