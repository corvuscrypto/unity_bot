package chat

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandMessage is just a specialized form of message that allows easier command handling
type CommandMessage struct {
	Type      string
	Arguments []string
	Message   *discordgo.MessageCreate
}

// CommandHandler represents a handler that takes a command name and possibly
// other positional arguments
type CommandHandler func(*discordgo.Session, *CommandMessage)

var commandHandlers = make(map[string]CommandHandler)

// AddCommandHandler allows handlers to be added on the fly
func AddCommandHandler(command string, handler CommandHandler) {
	commandHandlers[command] = handler
}

// RemoveHandler allows handlers to be removed on the fly
func RemoveHandler(command string) {
	delete(commandHandlers, command)
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	commandParts := strings.Split(strings.TrimSpace(m.Content[1:]), " ")
	command := new(CommandMessage)
	command.Type = commandParts[0]
	command.Message = m
	if len(commandParts) > 1 {
		command.Arguments = commandParts[1:]
	} else {
		command.Arguments = []string{}
	}
	if handler, ok := commandHandlers[command.Type]; ok {
		handler(s, command)
	}
}
