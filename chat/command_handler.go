package chat

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/corvuscrypto/unity_bot/config"
)

// CommandHandler represents a handler that takes a command name and possibly
// other positional arguments
type CommandHandler func(string, ...string)

var commandHandlers map[string]CommandHandler

// AddHandler allows handlers to be added on the fly
func AddHandler(command string, handler CommandHandler) {
	commandHandlers[command] = handler
}

// RemoveHandler allows handlers to be removed on the fly
func RemoveHandler(command string) {
	delete(commandHandlers, command)
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0] != config.GlobalConfig.CommandPrefix {
		return
	}

	commandParts := strings.Split(strings.TrimSpace(m.Content), " ")
	command := commandParts[0]
	if handler, ok := commandHandlers[command]; ok {
		if len(commandParts) > 1 {
			handler(command, commandParts[1:]...)
		} else {
			handler(command)
		}
	}
}
