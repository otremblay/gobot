package gobotplugin

import (
	"github.com/gabeguz/xmppbot"
	"math/rand"
	"strings"
	"time"
)

type Beer struct{}

func (p Beer) Name() string {
	return "Beer v1.0"
}

func (p Beer) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	// Don't reply if it's the bot asking
	if msg.From() == bot.FullName() {
		return nil
	}

	var reply string

	if strings.HasPrefix(msg.Body(), "beer?") {
		reply = beer()
	}

	if strings.HasPrefix(msg.Body(), "ビール?") {
		reply = beer()
	}

	if strings.HasPrefix(msg.Body(), "맥주?") {
		reply = beer()
	}

	bot.Send(reply)

	return nil
}

func beer() string {
	now := time.Now()
	if now.Weekday().String() == "Friday" && now.Hour() >= 16 {
		rand.Seed(time.Now().UnixNano())
		yup := ayes[rand.Intn(len(ayes))]
		return yup
	} else {
		rand.Seed(time.Now().UnixNano())
		nope := nays[rand.Intn(len(nays))]
		return nope
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

var ayes = []string{
	"proost! (beer)",
	"(beer) 乾杯",
	"cheers! (beer)",
	"¡salud! (beer)",
	"(beer) 건배",
}

var nays = []string{
	"wait for it...",
	"beer must wait.",
	"(yoda) Friday, the beer drinking day is. Until four o'clock must you wait.",
	"Really want to hit the ballmer peak eh?",
}
