package gobot

import (
	"log"
	"time"
)

type Bot interface {
	Name() string
	FullName() string
	Send(msg string)
	Connect() error
	PingServer(time.Duration)
	Listen() chan Message
	SetLogger(*log.Logger)
	Log(msg string)
}

//*************************************************
type Message interface {
	Body() string
	From() string
}

type Plugin interface {
	Name() string
	Execute(msg Message, bot Bot) error
}
