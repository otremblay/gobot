package slack

import (
	"fmt"
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
	client *slack.RTM
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
	b.logger.Printf("Connecting to slack\n")
	// slack.New doesn't return an error
	slack := slack.New(b.Opt.Token)
	b.client = slack.NewRTM()
	b.logger.Printf("Joining %s\n", b.Opt.Room)
	// TODO - join a room!
	return nil
}

// TODO this shouldn't be an exported part of the bot interface, as not all bots will need this.
func (b *bot) PingServer(seconds time.Duration) {
}

func (b *bot) Listen() chan gobot.Message {
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
				b.client.SendMessage(b.client.NewOutgoingMessage("Hello world", b.Opt.Room))

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				//recv <- message{body: ev}

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
				fmt.Printf("Other event: %v\n", ev)
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
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
