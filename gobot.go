package gobot

import (
	"log"
)

// Bot is the interface implemented by an object that can connect to a chat
// service and send and reply to messages.
type Bot interface {
	Name() string
	FullName() string
	Send(msg string)
	Reply(orig Message, msg string)
	Connect() error
	Listen() chan Message
	SetLogger(*log.Logger)
	Log(msg string)
}

// Message is the interface implemented by an object that contains an individual
// chat message as well as the meta-data associated with that chat message.
type Message interface {
	Body() string
	From() string
	Room() string
}

// Plugin is the interface implemented by an object that interacts with an
// individual chat message.
type Plugin interface {
	Name() string
	Execute(msg Message, bot Bot) error
}

// Gobot is a wrapper around a chatbot and its associated plugins.
type Gobot struct {
	Bot
	Plugins []Plugin
}

// InternalBot grants access to the currently used chatbot.
func (g *Gobot) InternalBot() Bot {
	return g.Bot
}
