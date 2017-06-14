package chat

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var channelsByGuild map[string][]*discordgo.Channel

func retrieveChannels() {
	channelsByGuild = make(map[string][]*discordgo.Channel)
	channelList, err := DiscordSession.UserChannels()
	if err != nil {
		log.Println(err)
		return
	}
	for _, channel := range channelList {
		channels, ok := channelsByGuild[channel.GuildID]
		if !ok {
			channelsByGuild[channel.GuildID] = make([]*discordgo.Channel, 0)
		}
		channelsByGuild[channel.GuildID] = append(channels, channel)
	}
}
