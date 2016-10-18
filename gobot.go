package main

import (
	"flag"
	"log"

	"github.com/gabeguz/gobotplugin/beer"
	"github.com/gabeguz/gobotplugin/chatlog"
	"github.com/gabeguz/gobotplugin/dm"
	"github.com/gabeguz/gobotplugin/echo"
	"github.com/gabeguz/gobotplugin/quote"
	"github.com/gabeguz/gobotplugin/rickroll"
	"github.com/gabeguz/gobotplugin/stathat"
	"github.com/gabeguz/gobotplugin/troll"
	"github.com/gabeguz/gobotplugin/url"
	"github.com/gabeguz/libgobot"
	"github.com/gabeguz/xmppbot"
)

func main() {
	var host, user, pass, room, name string
	flag.StringVar(&host, "host", "", "Hostname:port of the XMPP server")
	flag.StringVar(&user, "user", "", "Username of XMPP server (i.e.: foo@hostname.com")
	flag.StringVar(&pass, "pass", "", "Password for XMPP server")
	flag.StringVar(&room, "room", "", "Room to join (i.e.: #myroom@hostname.com")
	flag.StringVar(&name, "name", "gobot", "Name of the bot")
	flag.Parse()

	//TODO:Add some validation...but whatever for now

	bot := Gobot{
		xmppbot.New(host, user, pass, room, name),
		[]gobot.Plugin{
			echo.Echo{},
			beer.Beer{},
			quote.Quote{},
			dm.DirectMessage{},
			stathat.StatHat{},
			chatlog.ChatLog{},
			troll.Troll{},
			rickroll.RickRoll{},
			url.Url{},
		},
	}
	err := bot.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	go bot.PingServer(30)
	var msg gobot.Message
	var plugin gobot.Plugin
	for msg = range bot.Listen() {
		for _, plugin = range bot.Plugins {
			go executePlugin(plugin, msg, bot)
		}
	}

}

func executePlugin(p gobot.Plugin, m gobot.Message, b gobot.Bot) {
	err := p.Execute(m, b)
	if err != nil {
		b.Log(p.Name() + " => " + err.Error())
	}
}

type Gobot struct {
	gobot.Bot
	Plugins []gobot.Plugin
}
