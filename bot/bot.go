package GoBot

import (
	"fmt"
	"../config"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type GoBot struct {
	Token string
	Prefix string
	EmojiDir string

	eventsMutex    sync.Mutex
	currentEvents  []*CachedEvent
	upcomingEvents []*CachedEvent

	session *discordgo.Session
}

var (
	BotID string
	goBot *discordgo.Session
) 

func (g *GoBot) Start() {
	if g.session != nil {
		return
	}

	goBot, err := discordgo.New("Bot " + Config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(g.messageHandler)
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")

	g.session = goBot

	return
}

func (g *GoBot) messageHandler(s* discordgo.Session, m* discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	
	if Config.Prefix == "" {
		return
	}
	if !strings.HasPrefix(m.Content, Config.Prefix) {
		return
	}

	var (
		err error

		head string
		tail string
	)

	parts := strings.SplitN(m.Content, " ", 2)
	if len(parts) == 0 {
		return
	}

	head = strings.TrimPrefix(parts[0], Config.Prefix)
	if len(parts) > 1 {
		tail = parts[1]
	}

	fmt.Println("head=%#v tail=%#v", head, tail)

	switch head {
	case "events":
		err = g.cmdEvents(s, m)
	case "emo":
		err = g.cmdEmoji(s, m, tail)
	}

	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
	}
}