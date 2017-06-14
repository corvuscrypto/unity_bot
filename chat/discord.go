package chat

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/corvuscrypto/unity_bot/config"
)

// This file pertains to maintaining the discord connections

// DiscordSession is the main connection/session for discord (duh)
var DiscordSession *discordgo.Session

//ConnectToDiscord initiates a connection to discord.
func ConnectToDiscord() {
	var err error
	DiscordSession, err = discordgo.New("Bot " + config.GlobalConfig.Discord.Token)
	if err != nil {
		log.Fatal(err)
	}
	err = DiscordSession.Open()
	if err != nil {
		log.Fatal(err)
	}
	retrieveChannels()
	DiscordSession.AddHandler(handleMessageCreate)
}

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if it's a command, handle it
	if m.Content[0] == config.GlobalConfig.CommandPrefix {
		handleCommand(s, m)
		return
	}

	// otherwise treat it as a message
	handleMessage(s, m)
}
