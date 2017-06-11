package config

// This is where all the config defaults go for the bot.
const (
	DefaultBotName string = "Unity"
	DefaultUIPort  int    = 10101
)

func setDefaults() {
	GlobalConfig.BotName = DefaultBotName
	GlobalConfig.UIPort = DefaultUIPort

}
