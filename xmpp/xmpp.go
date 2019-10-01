package xmpp

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/gabeguz/gobot"
	"github.com/mattn/go-xmpp"
)

type Options struct {
	Room     string
	Protocol string
	Host     string
	User     string
	Password string
	Resource string
	NoTLS    bool
	StartTLS bool
	Debug    bool
	Session  bool
}

type message struct {
	body, from, room string
}

func (m message) Body() string {
	return m.body
}

func (m message) From() string {
	return m.from
}

func (m message) Room() string {
	return m.room
}

type Bot struct {
	Opt    Options
	client *xmpp.Client
	logger *log.Logger
}

func (b *Bot) FullName() string {
	return b.Opt.Room + "/" + b.Opt.Resource
}

func (b *Bot) Name() string {
	return b.Opt.Resource
}

func (b *Bot) Send(msg string) {
	b.client.Send(xmpp.Chat{Remote: b.Opt.Room, Type: "groupchat", Text: msg})
}

func (b *Bot) Reply(orig gobot.Message, msg string) {
	b.client.Send(xmpp.Chat{Remote: b.Opt.Room, Type: "groupchat", Text: msg})
}

func (b *Bot) Connect() error {
	var err error
	b.logger.Printf("Connecting to %s:*******@%s \n", b.Opt.User, b.Opt.Host)

	// Allow use of unknown certificates
	xmpp.DefaultConfig = tls.Config{
		InsecureSkipVerify: true,
	}

	options := xmpp.Options{
		Host:     b.Opt.Host,
		User:     b.Opt.User,
		Password: b.Opt.Password,
		NoTLS:    b.Opt.NoTLS,
		StartTLS: b.Opt.StartTLS,
		Debug:    b.Opt.Debug,
		Session:  b.Opt.Session,
	}
	b.client, err = options.NewClient()
	if err != nil {
		b.logger.Printf("Error: %s \n", err)
		return err
	}
	b.logger.Printf("Joining %s with resource %s \n", b.Opt.Room, b.Opt.Resource)
	b.client.JoinMUCNoHistory(b.Opt.Room, b.Opt.Resource)
	return nil
}

func (b *Bot) pingServer(seconds time.Duration) {
	if seconds > 0 {
		for _ = range time.Tick(seconds * time.Second) {
			b.client.PingC2S(b.Opt.Host+"/"+b.Opt.Resource, b.Opt.Host)
		}
	}
}

func (b *Bot) Listen() chan gobot.Message {
	msgChan := make(chan gobot.Message)

	go func(recv chan gobot.Message) {
		for {
			chat, err := b.client.Recv()
			if err != nil {
				b.logger.Printf("Error: %s \n", err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				recv <- message{body: v.Text, from: v.Remote, room: v.Remote}
			case xmpp.Presence:
				// ignore presence
			}
		}
	}(msgChan)

	return msgChan
}

func (b *Bot) SetLogger(logger *log.Logger) {
	b.logger = logger
}

func (b *Bot) Log(msg string) {
	b.logger.Printf("%s \n", msg)
}

// New returns a new XMPP Bot
func New(host, user, password, room, name string) *Bot {
	opt := Options{
		Host:     host,
		User:     user,
		Password: password,
		Resource: name,
		NoTLS:    true,
		StartTLS: true,
		Debug:    false,
		Session:  true,
		Room:     room,
		Protocol: "xmpp",
	}
	bot := &Bot{Opt: opt}
	bot.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	return bot
}
