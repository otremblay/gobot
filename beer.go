package seedbotplugin

import (
	"github.com/gabeguz/xmppbot"
	"strings"
	"time"
)

type Beer struct{}

func (p Beer) Name() string {
	return "Beer v1.0"
}

func (p Beer) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	if strings.HasPrefix(msg.Body(), "beer?") {
		if msg.From() != bot.FullName() {
			reply := beer()
			bot.Send(reply)
		}
	}

	if strings.HasPrefix(msg.Body(), "ビール?") {
		if msg.From() != bot.FullName() {
			reply := beerJapan()
			bot.Send(reply)
		}
	}

	return nil
}

func beer() string {
	now := time.Now()
	if now.Weekday().String() == "Friday" && now.Hour() >= 16 {
		return "proost! (beer)"
	} else {
		return "wait for it..."
	}
}

func beerJapan() string {
	now := time.Now()
	if now.Weekday().String() == "Friday" && now.Hour() >= 16 {
		return "(beer) おねがいします"
	} else {
		return "wait for it..."
	}
}
