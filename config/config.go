package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token                string
	OpenWeatherMapApiKey string
	ChannelID            string
	GuildID              string
	Config               *configStruct
)

type configStruct struct {
	Token                string `json:"Token"`
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
	ChannelID            string `json:"ChannelID"`
	GuildID              string `json:"GuildID"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := os.ReadFile(".config")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(file)

	err = json.Unmarshal(file, &Config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = Config.Token
	OpenWeatherMapApiKey = Config.OpenWeatherMapApiKey
	ChannelID = Config.ChannelID
	GuildID = Config.GuildID

	return nil
}
