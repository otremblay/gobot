package url

import (
	"fmt"
	"github.com/gabeguz/libgobot"
	"github.com/thatguystone/swan"
	"net/url"
)

type Url struct{}

func (p Url) Name() string {
	return "Url v1.0"
}

func (p Url) Execute(msg gobot.Message, bot gobot.Bot) error {

	if msg.From() == bot.FullName() {
		return nil
	}

	u, err := url.ParseRequestURI(msg.Body())
	if err != nil {
		return err
	}

	a, err := swan.FromURL(u.String())
	if err != nil {
		return (err)
	}

	// Respond with the article title
	fmt.Printf("Title: %v\n", a)
	bot.Send(a.Meta.Title)

	return nil
}
