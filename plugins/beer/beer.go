package beer

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gabeguz/gobot/bot"
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

func (p Beer) Execute(msg bot.Message, bot bot.Bot) error {
	// Don't reply if it's the bot asking
	if msg.From() == bot.FullName() {
		return nil
	}

	now := time.Now()
	if strings.HasPrefix(msg.Body(), "beer?") {
		bot.Reply(msg, beer(now))
	} else if strings.HasPrefix(msg.Body(), "ビール?") {
		bot.Reply(msg, beer(now))
	} else if strings.HasPrefix(msg.Body(), "맥주?") {
		bot.Reply(msg, beer(now))
	}

	return nil
}

func beer(t time.Time) string {
	if t.Weekday().String() == "Friday" && t.Hour() >= 16 {
		rand.Seed(time.Now().UnixNano())
		yup := ayes[rand.Intn(len(ayes))]
		return yup
	} else if t.Month().String() == "December" && t.Day() == 24 {
		return "Merry Beermas!"
	} else {
		rand.Seed(time.Now().UnixNano())
		nope := nays[rand.Intn(len(nays))]
		return nope
	}
}
