package config

// This is where all the config defaults go for the bot.
const (
	DefaultBotName string = "Unity"
	DefaultUIPort int = 10101
	DefaultCommandPrefix byte = '/'

	// DB constants
	DefaultDBUser string = ""
	DefaultDBPassword string = ""
	DefaultDBAddress string = ""
	DefaultDBPort int = 3306
	DefaultDBName string = "unity_db"
)

func setDefaults() {
	GlobalConfig.BotName = DefaultBotName
	GlobalConfig.UIPort = DefaultUIPort
	GlobalConfig.CommandPrefix = DefaultCommandPrefix

	GlobalConfig.DB.User = DefaultDBUser
	GlobalConfig.DB.Password = DefaultDBPassword
	GlobalConfig.DB.Address = DefaultDBAddress
	GlobalConfig.DB.Port = DefaultDBPort
	GlobalConfig.DB.Name = DefaultDBName
}
