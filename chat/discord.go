package chat

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/corvuscrypto/unity_bot/config"
)

// This file pertains to maintaining the discord connections

// DiscordSession is the main connection/session for discord (duh)
var DiscordSession *discordgo.Session

func connectToDiscord() {
	var err error
	DiscordSession, err = discordgo.New("bot " + config.GlobalConfig.Discord.Token)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	connectToDiscord()
	retrieveChannels()
}
