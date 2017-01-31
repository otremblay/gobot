package xmpp

import (
	"github.com/gabeguz/gobot"
	"github.com/nlopes/slack"
	"log"
	"os"
	"time"
)

type Options struct {
	Token string
	Room  string
	Name  string
	Debug bool
}

//*************************************************
type message struct {
	body, from string
}

func (m message) Body() string {
	return m.body
}

func (m message) From() string {
	return m.from
}

//*************************************************
type bot struct {
	Opt    Options
	client *slack.Client
	logger *log.Logger
}

func (b *bot) FullName() string {
	return b.Opt.Name
}

func (b *bot) Name() string {
	return b.Opt.Name
}

func (b *bot) Send(msg string) {
}

func (b *bot) Connect() error {
	var err error
	b.logger.Printf("Connecting to slack\n")
	// slack.New doesn't return an error
	b.client = slack.New(b.Opt.Token)
	b.logger.Printf("Joining %s\n", b.Opt.Room)
	// TODO - join a room!
	return nil
}

// TODO this shouldn't be an exported part of the bot interface, as not all bots will need this.
func (b *bot) PingServer(seconds time.Duration) {
}

func (b *bot) Listen() chan gobot.Message {
	msgChan := make(chan gobot.Message)

	go func(recv chan gobot.Message) {
		for {
			chat, err := b.client.Recv()
			if err != nil {
				b.logger.Printf("Error: %s \n", err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				recv <- message{body: v.Text, from: v.Remote}
			case xmpp.Presence:
				b.logger.Printf("Presence: %+v \n", v)
			}
		}
	}(msgChan)

	return msgChan
}

func (b *bot) SetLogger(logger *log.Logger) {
	b.logger = logger
}

func (b *bot) Log(msg string) {
	b.logger.Printf("%s \n", msg)
}

//*************************************************

func New(token, room, name string) gobot.Bot {
	opt := Options{
		Token: token,
		Room:  room,
		Name:  name,
		Debug: true,
	}
	bot := &bot{Opt: opt}
	bot.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	return bot
}
