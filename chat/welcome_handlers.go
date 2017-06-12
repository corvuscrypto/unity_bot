package chat

import (
  "fmt"

  "github.com/bwmarrin/discordgo"
)


func memberWelcome(s *discordgo.Session, evt *discordgo.GuildMemberAdd) {
  if evt.GuildID != "certainid" {
    return
  }
  message := fmt.Sprintf("Please welcome @%s", evt.Nick)
  s.ChannelMessageSend("", message)
}

func memberLeave(s *discordgo.Session, evt *discordgo.GuildMemberRemove) {
  if evt.GuildID != "certainid" {
    return
  }
  message := fmt.Sprintf("Unfortunately, %s has left the group.", evt.Nick)
  s.ChannelMessageSend("", message)
}


func addWelcomeHandlers() {
	DiscordSession.AddHandler(memberWelcome)
	DiscordSession.AddHandler(memberLeave)
}
