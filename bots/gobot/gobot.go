package gobot

import (
	"github.com/gabeguz/gobot"
)

type Gobot struct {
	gobot.Bot
	Plugins []gobot.Plugin
}

func (g *Gobot) InternalBot() gobot.Bot {
	return g.Bot
}
