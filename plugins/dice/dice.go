package dice

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"math/rand"
	"time"

	"github.com/gabeguz/gobot"
	gb "github.com/gabeguz/gobot/bots/gobot"
	sb "github.com/gabeguz/gobot/bots/slack"
	"github.com/gabeguz/gobot/plugins/chatlog"
	"github.com/nlopes/slack"
)

type Dice struct {
	Filename string
	Log      chatlog.ChatLog
}

var diceRE = regexp.MustCompile(`^roll ([0-9]{1,3})d([0-9]{1,3}) ?([+-][0-9]+)?`)

func (Dice) Name() string {
	return "Dice v1.0"
}

var users = map[string]string{}

func (d Dice) Execute(msg gobot.Message, bot gobot.Bot) error {
	b2 := bot.(gb.Gobot)
	if msg.From() != bot.FullName() {
		if matches := diceRE.FindAllStringSubmatch(msg.Body(), -1); len(matches) > 0 {
			rand.Seed(time.Now().UnixNano())
			var rollresult = []int{}
			dice, _ := strconv.Atoi(matches[0][2])
			total := 0
			bonus := 0
			if len(matches[0]) > 2 {
				bonus, _ = strconv.Atoi(strings.TrimSpace(matches[0][3]))
			}
			for i, _ := strconv.Atoi(matches[0][1]); i > 0; i -= 1 {
				number := rand.Intn(dice) + 1
				total += number
				rollresult = append(rollresult, number)
			}
			switch b3 := b2.InternalBot().(type) {
			case *sb.Bot:
				tmpmsg := msg.(sb.Message)
				mess := tmpmsg.EffectiveMessage()

				c := b3.Client()
				p := slack.NewPostMessageParameters()
				p.EscapeText = false
				p.Username = b3.Opt.Name
				var n string
				if user, ok := users[mess.User]; ok {
					n = user
				} else {
					u, err := b3.Client().GetUserInfo(mess.User)
					if err == nil {

						n = u.RealName
						users[mess.User] = n
					}
				}
				result := ""
				if bonus == 0 {
					result = fmt.Sprintf("%s rolls %sd%s for a total of %d", n, matches[0][1], matches[0][2], total)
				} else {
					result = fmt.Sprintf("%s rolls %sd%s for a total of %d (%s) = %d", n, matches[0][1], matches[0][2], total, matches[0][3], total+bonus)
				}
				d.Log.Logit(n, result)
				p.Attachments = []slack.Attachment{
					slack.Attachment{
						Title: result,
						Text:  fmt.Sprint(rollresult, matches[0][3]),
					},
				}

				if mess.ThreadTimestamp != "" {
					p.ThreadTimestamp = mess.ThreadTimestamp
				}
				c.PostMessage(mess.Channel, "", p)
			default:
				b := bytes.NewBuffer(nil)
				fmt.Fprintln(b, fmt.Sprintf("%s rolls %sd%s: %v", msg.From(), matches[0][1], matches[0][2], rollresult))
				d.Log.Logit(msg.From(), b.String())
				bot.Send(b.String())
			}
		}
	}
	return nil
}
