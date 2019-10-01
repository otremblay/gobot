package beer

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gabeguz/gobot"
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
	"Not time for beer?  Try a scotch.",
}

var scotchAyes = []string{
	"Mmm... scotch.",
	"Scotch, scotch, scotch, I love scotch",
	"Feeling peaty are we? Cheers.",
}

var scotchNays = []string{
	"Sorry, I just poured myself the last glass.",
	"While I'm never opposed to a glass of scotch, your coworkers might not appreciate you having one right *now.*",
	"no scotch for you.",
}

type Beer struct{}

func (p Beer) Name() string {
	return "Beer v1.0"
}

func (p Beer) Execute(msg gobot.Message, bot gobot.Bot) error {
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
	} else if strings.HasPrefix(msg.Body(), "scotch?") {
		bot.Reply(msg, scotch(now))
	}

	return nil
}

func beer(t time.Time) string {
	var msg string
	if t.Weekday().String() == "Friday" && t.Hour() >= 16 {
		rand.Seed(time.Now().UnixNano())
		msg = ayes[rand.Intn(len(ayes))]
	} else if t.Month().String() == "December" && t.Day() == 24 {
		msg = "Merry Beermas!"
	} else {
		rand.Seed(time.Now().UnixNano())
		msg = nays[rand.Intn(len(nays))]
	}
	return msg
}

func scotch(t time.Time) string {
	var msg string
	if t.Hour() >= 14 {
		rand.Seed(time.Now().UnixNano())
		msg = scotchAyes[rand.Intn(len(scotchAyes))]
	} else {
		rand.Seed(time.Now().UnixNano())
		msg = scotchNays[rand.Intn(len(scotchNays))]
	}
	return msg
}
