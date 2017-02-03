package echo

import (
	"github.com/gabeguz/libgobot"
	"log"
	"testing"
	"time"
)

type Bot struct{}

func (p Bot) Connect() error {
	return nil
}

func (p Bot) Listen() chan gobot.Message {
	return nil
}

func (p Bot) Log(string) {
}

func (p Bot) Name() string {
	return "bot"
}

func (p Bot) FullName() string {
	return "testbot"
}

func (p Bot) Send(string) {
}

func (p Bot) SetLogger(*log.Logger) {
}

func (p Bot) PingServer(time.Duration) {
}

type Message struct{}

func (m Message) Body() string {
	return "this is a test of the emergency broadcast network."
}

func (m Message) From() string {
	return "nottestbot"
}

func TestPluginName(t *testing.T) {
	p := Echo{}
	expected := "Echo v1.0"
	name := p.Name()
	if name != expected {
		t.Errorf("Test: '%s' does not equal '%s'", name, expected)
	}

}

func TestEcho(t *testing.T) {
	message := Message{}
	bot := Bot{}
	p := Echo{}
	err := p.Execute(message, bot)
	if err != nil {
		t.Errorf("Error sending message")
	}
}
