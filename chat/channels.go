package chat

import "github.com/bwmarrin/discordgo"

var channelsByGuild map[string][]*discordgo.Channel

func retrieveChannels() {
  channelsByGuild = make(map[string][]*discordgo.Channel)
  for _, guildID := range []string{"ID1", "ID2", "ID3"} {
    channels, _ := DiscordSession.GuildChannels(guildID)
    channelsByGuild[guildID] = channels
  }
}
