package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gabeguz/gobot"
	"github.com/gabeguz/gobot/bot"
	gb "github.com/gabeguz/gobot/bots/gobot"
	"github.com/gabeguz/gobot/bots/slack"
	"github.com/gabeguz/gobot/plugins/beer"
	"github.com/gabeguz/gobot/plugins/chatlog"
	"github.com/gabeguz/gobot/plugins/cron"
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
	var host, user, pass, room, name string
	var crons strslice
	flag.StringVar(&host, "host", "", "Hostname:port of the XMPP server")
	flag.StringVar(&user, "user", "", "Username of XMPP server (i.e.: foo@hostname.com")
	flag.StringVar(&pass, "pass", "", "Password for XMPP server")
	flag.StringVar(&room, "room", "", "Room to join (i.e.: #myroom@hostname.com")
	flag.StringVar(&name, "name", "gobot", "Name of the bot")
	flag.Var(&crons, "job", "List of jobs")
	flag.Parse()

	//TODO:Add some validation...but whatever for now

	bot := gb.Gobot{
		slack.New(pass, room, name),
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
	var msg bot.Message
	var plugin bot.Plugin
	for msg = range gobot.Listen() {
		for _, plugin = range gobot.plugins {
			go executePlugin(plugin, msg, gobot)
		}
	}

}

func executePlugin(p bot.Plugin, m bot.Message, b bot.Bot) {
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

