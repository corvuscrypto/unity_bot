package chat

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

// MessageHandler is a function that is needed separate of discordgo for multiplexing
type MessageHandler func(*discordgo.Session, *discordgo.MessageCreate)

// RegExMux allows traversal of an array to find a match for messages for multiplexing
type RegExMux struct {
	RegEx   *regexp.Regexp
	Handler MessageHandler
}

var muxArray = make([]RegExMux, 0)

// AddMessageHandler adds an instance of RegExMux to the internal array
func AddMessageHandler(r *regexp.Regexp, handler MessageHandler) {
	muxArray = append(muxArray, RegExMux{
		r,
		handler,
	})
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, muxer := range muxArray {
		if muxer.RegEx.MatchString(m.Content) {
			muxer.Handler(s, m)
			break
		}
	}
}
