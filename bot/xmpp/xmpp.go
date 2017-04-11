package xmpp

import (
	"log"
	"os"
	"time"

	"github.com/gabeguz/gobot/bot"
	"github.com/mattn/go-xmpp"
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
type xmppbot struct {
	Opt    Options
	client *xmpp.Client
	logger *log.Logger
}

func (b *xmppbot) FullName() string {
	return b.Opt.Room + "/" + b.Opt.Resource
}

func (b *xmppbot) Name() string {
	return b.Opt.Resource
}

func (b *xmppbot) Send(msg string) {
	b.client.Send(xmpp.Chat{Remote: b.Opt.Room, Type: "groupchat", Text: msg})
}

func (b *xmppbot) Connect() error {
	var err error
	b.logger.Printf("Connecting to %s:*******@%s \n", b.Opt.User, b.Opt.Host)
	b.client, err = b.Opt.NewClient()
	if err != nil {
		b.logger.Printf("Error: %s \n", err)
		return err
	}
	b.logger.Printf("Joining %s with resource %s \n", b.Opt.Room, b.Opt.Resource)
	b.client.JoinMUCNoHistory(b.Opt.Room, b.Opt.Resource)
	return nil
}

func (b *xmppbot) pingServer(seconds time.Duration) {
	if seconds > 0 {
		for _ = range time.Tick(seconds * time.Second) {
			b.client.PingC2S(b.Opt.Host+"/"+b.Opt.Resource, b.Opt.Host)
		}
	}
}

func (b *xmppbot) Listen() chan bot.Message {
	msgChan := make(chan bot.Message)

	go func(recv chan bot.Message) {
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

func (b *xmppbot) SetLogger(logger *log.Logger) {
	b.logger = logger
}

func (b *xmppbot) Log(msg string) {
	b.logger.Printf("%s \n", msg)
}

//*************************************************

func New(host, user, password, room, name string) bot.Bot {
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
	xmppbot := &xmppbot{Opt: opt}
	xmppbot.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	return xmppbot
}
