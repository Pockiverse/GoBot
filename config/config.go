package Config

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

var (
	Token string
	Prefix string
	EmojiDir string

	config *configStruct
)

type configStruct struct {
	Token string `json:"token"`
	Prefix string `json:"prefix"`
	EmojiDir string `json:"emoji_dir"`
}

func LoadConfigFile() error {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	Prefix = config.Prefix
	EmojiDir = config.EmojiDir

	return nil
}
