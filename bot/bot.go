package bot

import (
	"fmt"
	"strconv"

	"github.com/agusheryanto182/go-bot-discord-orders/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotID     string
	ChannelID string
	GuildID   string

	goBot *discordgo.Session
)

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "who" {
		s.ChannelMessageSend(m.ChannelID, "i am suga waifu")
	} else if m.Content == "help" {
		helpMessage := "- who : tell about me\n" +
			"- help : see the command\n" +
			"- earthquake information : see the latest earthquake information \n" +
			"- write a place if you want to know the weather temperature : ex -> Sukabumi\n"

		s.ChannelMessageSend(m.ChannelID, helpMessage)
	} else if m.Content != "" && m.Content != "help" && m.Content != "who" && m.Content != "earthquake information" {
		data, err := query(m.Content)
		if err != nil {
			fmt.Println(err)
		}

		name := string(data.Name)

		celcius := data.Main.Kelvin - 273.15

		celciusStr := strconv.FormatFloat(celcius, 'f', 0, 64)

		message := "Lokasi : " + name + "\n" + "Temperatur : " + celciusStr + " °C" + "\n"

		s.ChannelMessageSend(m.ChannelID, message)
	} else if m.Content == "earthquake information" {
		data, err := getGempa()
		if err != nil {
			fmt.Println(err.Error())
		}

		s.ChannelMessageSend(m.ChannelID,
			"Wilayah : "+data.Infogempa.Gempa.Wilayah+"\n"+
				"Jam : "+data.Infogempa.Gempa.Jam+"\n"+
				"Tanggal : "+data.Infogempa.Gempa.Tanggal+"\n"+
				"Datetime : "+data.Infogempa.Gempa.DateTime+"\n"+
				"Koordinat : "+data.Infogempa.Gempa.Coordinates+"\n"+
				"Kedalaman : "+data.Infogempa.Gempa.Kedalaman+"\n"+
				"Potensi : "+data.Infogempa.Gempa.Potensi+"\n"+
				"Dirasakan : "+data.Infogempa.Gempa.Dirasakan+"\n")
	}
}
