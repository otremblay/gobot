package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	gb "github.com/gabeguz/gobot"
	"github.com/gabeguz/gobot/plugins/beer"
	"github.com/gabeguz/gobot/plugins/chatlog"
	"github.com/gabeguz/gobot/plugins/cron"
	"github.com/gabeguz/gobot/plugins/dice"
	"github.com/gabeguz/gobot/plugins/echo"
	"github.com/gabeguz/gobot/plugins/jira"
	"github.com/gabeguz/gobot/plugins/quote"
	"github.com/gabeguz/gobot/slack"
	"github.com/gabeguz/gobot/xmpp"
)

var host, user, pass, room, name, protocol, logfile string
var crons strslice

func init() {
	flag.StringVar(&host, "host", "", "Hostname:port of the XMPP server")
	flag.StringVar(&user, "user", "", "Username of XMPP server (i.e.: foo@hostname.com")
	flag.StringVar(&pass, "pass", "", "Password for XMPP server")
	flag.StringVar(&room, "room", "", "Room to join (i.e.: #myroom@hostname.com")
	flag.StringVar(&name, "name", "gobot", "Name of the bot")
	flag.StringVar(&protocol, "protocol", "xmpp", "Protocol (xmpp, slack)")
	flag.StringVar(&logfile, "logfile", "/tmp/chatlog", "Path to log file")
	flag.Var(&crons, "job", "List of jobs")
}

func createBot(plugins []gb.Plugin) gb.Gobot {
	var bot gb.Gobot
	if protocol == "slack" {
		bot = gb.Gobot{
			slack.New(pass, room, name),
			plugins,
		}
	} else {
		bot = gb.Gobot{
			xmpp.New(host, user, pass, room, name),
			plugins,
		}
	}
	return bot
}

func executePlugin(p gb.Plugin, m gb.Message, b gb.Bot) {
	err := p.Execute(m, b)
	if err != nil {
		b.Log(p.Name() + " => " + err.Error())
	}
}

func main() {
	flag.Parse()
	chatlog := chatlog.ChatLog{Filename: logfile}

	plugins := []gb.Plugin{
		echo.Echo{},
		beer.Beer{},
		quote.Quote{},
		dice.Dice{Log: chatlog},
		chatlog,
		jira.Jira{},
	}

	bot := createBot(plugins)
	err := bot.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	for _, crn := range crons {
		parts := strings.Split(crn, "|")
		cron.NewCron(parts[2], crn, bot)
	}

	var msg gb.Message
	var plugin gb.Plugin
	for msg = range bot.Listen() {
		for _, plugin = range bot.Plugins {
			go executePlugin(plugin, msg, bot)
		}
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
