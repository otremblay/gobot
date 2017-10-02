package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gabeguz/gobot"
	gb "github.com/gabeguz/gobot/bots/gobot"
	"github.com/gabeguz/gobot/bots/slack"
	"github.com/gabeguz/gobot/plugins/beer"
	"github.com/gabeguz/gobot/plugins/chatlog"
	"github.com/gabeguz/gobot/plugins/cron"
	"github.com/gabeguz/gobot/plugins/dice"
	"github.com/gabeguz/gobot/plugins/dm"
	"github.com/gabeguz/gobot/plugins/echo"
	"github.com/gabeguz/gobot/plugins/jira"
	"github.com/gabeguz/gobot/plugins/quote"
	"github.com/gabeguz/gobot/plugins/rickroll"
	"github.com/gabeguz/gobot/plugins/stathat"
	"github.com/gabeguz/gobot/plugins/troll"
	"github.com/gabeguz/gobot/plugins/url"
)

func main() {
	var host, user, pass, room, name, logfile string
	var crons strslice
	flag.StringVar(&host, "host", "", "Hostname:port of the XMPP server")
	flag.StringVar(&user, "user", "", "Username of XMPP server (i.e.: foo@hostname.com")
	flag.StringVar(&pass, "pass", "", "Password for XMPP server")
	flag.StringVar(&room, "room", "", "Room to join (i.e.: #myroom@hostname.com")
	flag.StringVar(&name, "name", "gobot", "Name of the bot")
	flag.StringVar(&logfile, "logfile", "/tmp/chatlog", "Name of the bot")
	flag.Var(&crons, "job", "List of jobs")
	flag.Parse()

	//TODO:Add some validation...but whatever for now
	chatlog := chatlog.ChatLog{Filename: logfile}

	bot := gb.Gobot{
		slack.New(pass, room, name),
		[]gobot.Plugin{
			echo.Echo{},
			beer.Beer{},
			quote.Quote{},
			dm.DirectMessage{},
			dice.Dice{Log: chatlog},
			chatlog,
			stathat.StatHat{},
			troll.Troll{},
			rickroll.RickRoll{},
			url.Url{},
			jira.Jira{},
		},
	}
	/*
		bot := gb.Gobot{
			xmpp.New(host, user, pass, room, name),
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
	*/

	err := bot.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	for _, crn := range crons {
		parts := strings.Split(crn, "|")
		cron.NewCron(parts[2], crn, bot)
	}

	//go bot.PingServer(30)
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

type strslice []string

func (s *strslice) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *strslice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
