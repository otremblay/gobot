package xmppbot

import (
	"github.com/gabeguz/gobot"
	"github.com/mattn/go-xmpp"
	"log"
	"os"
	"time"
)

type Options struct {
	xmpp.Options
	Room string
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
	client *xmpp.Client
	logger *log.Logger
}

func (b *bot) FullName() string {
	return b.Opt.Room + "/" + b.Opt.Resource
}

func (b *bot) Name() string {
	return b.Opt.Resource
}

func (b *bot) Send(msg string) {
	b.client.Send(xmpp.Chat{Remote: b.Opt.Room, Type: "groupchat", Text: msg})
}

func (b *bot) Connect() error {
	var err error
	b.logger.Printf("Connecting to %s:*******@%s \n", b.Opt.User, b.Opt.Host)
	b.client, err = b.Opt.NewClient()
	if err != nil {
		b.logger.Printf("Error: %s \n", err)
		return err
	}
	b.logger.Printf("Joining %s with resource %s \n", b.Opt.Room, b.Opt.Resource)
	b.client.JoinMUC(b.Opt.Room, b.Opt.Resource)
	return nil
}

func (b *bot) PingServer(seconds time.Duration) {
	if seconds > 0 {
		for _ = range time.Tick(seconds * time.Second) {
			b.client.PingC2S(b.Opt.Host+"/"+b.Opt.Resource, b.Opt.Host)
		}
	}
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

func New(host, user, password, room, name string) gobot.Bot {
	opt := Options{
		xmpp.Options{
			Host:     host,
			User:     user,
			Password: password,
			Resource: name,
			NoTLS:    true,
			Debug:    true,
			Session:  true,
		},
		room,
	}
	bot := &bot{Opt: opt}
	bot.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	return bot
}
