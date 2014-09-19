package main

import (
	"flag"
	"github.com/seedboxtech/seedbotplugin"
	"github.com/seedboxtech/xmppbot"
	"log"
)

func main() {
	var host, user, pass, room, name string
	flag.StringVar(&host, "host", "", "Hostname:port of the XMPP server")
	flag.StringVar(&user, "user", "", "Username of XMPP server (i.e.: foo@hostname.com")
	flag.StringVar(&pass, "pass", "", "Password for XMPP server")
	flag.StringVar(&room, "room", "", "Room to join (i.e.: #myroom@hostname.com")
	flag.StringVar(&name, "name", "seedbot", "Name of the bot")
	flag.Parse()

	//TODO:Add some validation...but whatever for now

	bot := Seedbot{
		xmppbot.New(host, user, pass, room, name),
		[]seedbotplugin.Plugin{
			seedbotplugin.Echo{},
			seedbotplugin.Quote{},
			seedbotplugin.DirectMessage{},
			seedbotplugin.StatHat{},
			seedbotplugin.ChatLog{},
			seedbotplugin.Jira{},
			seedbotplugin.Troll{},
		},
	}
	err := bot.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	for msg := range bot.Listen() {
		for _, p := range bot.Plugins {
			err := p.Execute(msg, bot)
			if err != nil {
				bot.Log(p.Name() + " => " + err.Error())
			}
		}
	}
}

type Seedbot struct {
	xmppbot.Bot
	Plugins []seedbotplugin.Plugin
}
