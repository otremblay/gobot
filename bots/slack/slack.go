package slack

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gabeguz/gobot"
	"github.com/nlopes/slack"
)

type Options struct {
	Token string
	Room  string
	Name  string
	Debug bool
}

//*************************************************
type Message struct {
	body, from       string
	effectiveMessage *slack.MessageEvent
}

func (m Message) Body() string {
	return m.body
}

func (m Message) From() string {
	return m.from
}

func (m Message) Room() string {
	return m.effectiveMessage.Channel
}

func (m Message) EffectiveMessage() *slack.MessageEvent {
	return m.effectiveMessage
}

//*************************************************
type Bot struct {
	Opt    Options
	client *slack.RTM
	logger *log.Logger
}

func (b *Bot) FullName() string {
	return b.Opt.Name
}

func (b *Bot) Name() string {
	return b.Opt.Name
}

func (b *Bot) Send(msg string) {
	b.client.SendMessage(b.client.NewOutgoingMessage(msg, b.Opt.Room))
}

func (b *Bot) Reply(orig gobot.Message, msg string) {
	b.client.SendMessage(b.client.NewOutgoingMessage(msg, orig.Room()))
}

func (b *Bot) Client() *slack.RTM {
	return b.client
}

func (b *Bot) Connect() error {
	b.logger.Printf("Connecting to slack\n")
	// slack.New doesn't return an error
	slack := slack.New(b.Opt.Token)
	slack.SetDebug(true)
	b.client = slack.NewRTM()
	b.logger.Printf("Joining %s\n", b.Opt.Room)
	return nil
}

// TODO this shouldn't be an exported part of the Bot interface, as not all Bots will need this.
func (b *Bot) PingServer(seconds time.Duration) {
}

func (b *Bot) Listen() chan gobot.Message {
	msgChan := make(chan gobot.Message)

	go b.client.ManageConnection()

	go func(recv chan gobot.Message) {
		for msg := range b.client.IncomingEvents {
			fmt.Print("Event Received: ")

			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello

			case *slack.ConnectedEvent:
				fmt.Println("Infos:", ev.Info)
				fmt.Println("Connection counter:", ev.ConnectionCount)
				// Replace #general with your Channel ID
				// b.client.SendMessage(b.client.NewOutgoingMessage("Hello world", b.Opt.Room))

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				recv <- Message{body: ev.Msg.Text, from: ev.Msg.Channel, effectiveMessage: ev}
				// recv <- Message{body: v.Text, from: v.Remote}

			case *slack.PresenceChangeEvent:
				fmt.Printf("Presence Change: %v\n", ev)
				b.logger.Printf("Presence Change: %+v \n", ev)

			case *slack.LatencyReport:
				fmt.Printf("Current latency: %v\n", ev.Value)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				return

			default:
				//fmt.Printf("Other event: %v\n", ev)
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
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

//*************************************************

func New(token, room, name string) gobot.Bot {
	opt := Options{
		Token: token,
		Room:  room,
		Name:  name,
		Debug: true,
	}
	Bot := &Bot{Opt: opt}
	Bot.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	return Bot
}
