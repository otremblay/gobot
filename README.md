gobot
=====

gobot is a Simple chatbot written in golang  

Initail work based on: https://github.com/mattn/go-xmpp/blob/master/example/example.go

gobot is under active development, the api may change w/out warning.  

By itself, gobot just sits and listens for incoming messages.  When a
message is received, gobot passes the message on to any number of plugins.  The
plugins provide all the features of the chatbot.

Installing gobot
----------------
First, you'll need to have golang installed:  http://golang.org/doc/install

Then, once that is working, all you need to do is: 

```
$ go get github.com/gabeguz/gobot/cmd/gobot
```

Running a bot
-------------
gobot supports both slack and XMPP (jabber).

To connect to slack you'll need an API token, directions for getting one can be
found at https://api.slack.com/bot-users aside from that, pick a name for your
bot and enter the room you would like your bot to join.

```
$ gobot -protocol=slack -user=herman -pass=<slack api token> -room=bottest -name=herman
```

For xmpp you need a few extra bits of information.  The hostname with port, the
abber user, the users password, and the full address to the MUC you'd like the
bot to join: 

```
$ gobot -protocol=xmpp -host="jabber.my.domain:5222" -user="gabe@jabber.my.domain" -pass="hunter2" -room="#gobottest@conference.jabber.my.domain" -name="gobot"
```

Creating A Plugin
-----------------

Creating a plugin is pretty easy, you just need to import the gobot library: 

```
import (
	"github.com/gabeguz/gobot"
)
```

And then implement 2 exported methods so that you conform to the bot.Plugin interface: 

```
type MyPlugin struct {}

func (p MyPlugin) Name() string {
	return "MyPlugin v1.0"
}

func (p MyPlugin) Execute(msg bot.Message, bot bot.Bot) error {
	if msg.From() != bot.FullName() {
		bot.Send("Why, hello!")
	}
	return nil
}
```

One last step, edit gobot/cmd/main.go and add a call to your new plugin
to the []gb.Plugin slice:

```
plugins := []gb.Plugin{
	myplugin.MyPlugin{},
}
```
and import the code in the `import` statement at the top of the file.

Finally, reinstall gobot `$ cd $GOPATH/src/github/gabeguz/gobot/cmd && go
install ./...`
