package gobotplugin

import (
	"github.com/gabeguz/xmppbot"
)

type Plugin interface {
	Name() string
	Execute(msg xmppbot.Message, bot xmppbot.Bot) error
}
