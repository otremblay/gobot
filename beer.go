package gobotplugin

import (
	"github.com/gabeguz/xmppbot"
	"math/rand"
	"strings"
	"time"
)

var ayes = []string{
	"proost! (beer)",
	"(beer) 乾杯",
	"cheers! (beer)",
	"¡salud! (beer)",
	"(beer) 건배",
	"Mmm... beer. (drool)",
}

var nays = []string{
	"wait for it...",
	"beer must wait.",
	"(yoda) Friday, the beer drinking day is. Until four o'clock must you wait.",
	"Really want to hit the ballmer peak eh?",
	"no beer for you.",
}

type Beer struct{}

func (p Beer) Name() string {
	return "Beer v1.0"
}

func (p Beer) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	// Don't reply if it's the bot asking
	if msg.From() == bot.FullName() {
		return nil
	}

	if strings.HasPrefix(msg.Body(), "beer?") {
		bot.Send(beer())
	} else if strings.HasPrefix(msg.Body(), "ビール?") {
		bot.Send(beer())
	} else if strings.HasPrefix(msg.Body(), "맥주?") {
		bot.Send(beer())
	}

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
