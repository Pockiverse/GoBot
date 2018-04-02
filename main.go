package main

import (
	"fmt"
	"net/http"
    "os"
    "log"

	"github.com/evanstan/GoBot/bot"
	"github.com/evanstan/GoBot/config"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
        return
    }

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

	http.ListenAndServe(":"+port, nil)
}