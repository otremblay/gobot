// Package bot provides the interfaces for bots, messages, and plugins used to
// implement a chatbot.
package bot

import (
	"log"
)

// Bot is the interface that must be implemented by a chat bot.  For example, an
// xmpp bot, a slack bot, an irc bot, etc.
type Bot interface {
	Name() string
	FullName() string
	Send(msg string)
	Connect() error
	Listen() chan Message
	SetLogger(*log.Logger)
	Log(msg string)
}

// Message is a message from the Bot that gets forwarded on to the Plugins for
// action to be taken.
type Message interface {
	Body() string
	From() string
}

// Plugin defines the interface that a valid Plugin must conform to in order to
// interact with chat bots.
type Plugin interface {
	Name() string
	Execute(msg Message, bot Bot) error
}
