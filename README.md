gobot
=====

gobot is a Simple XMPP bot written in golang  

Initail work based on https://github.com/mattn/go-xmpp/blob/master/example/example.go

gobot is under active development, the api may change w/out warning.  

By itself, gobot just sits and listens for incoming XMPP messages.  When a
message is received, gobot passes the message on to any number of plugins.  The
plugins provide all the features of the chatbot.

Currently, the following plugins exist: 

1. dm - gobot recognizes his nick and replies that his ears are burning when he is mentioned. 

2. quote - gobot understands the quote command.  Currently 3 types of quotes are supported: 
	code quote - quotes about programming
	integration quote - quotes about front end integration
	admin quote - quotes about systems administration

3. chatlog - gobot logs all messages to a text file. Currently '/tmp/chatlog'

4. stathat - gobot counts each message, and logs it as a single 'hit' at stathat.

5. keyword - gobot scans a message for work counts

6. echo - gobot echoes a message back to the room

7. beer - ask gobot for a beer?

Installing gobot
----------------
First, you'll need to have golang installed:  http://golang.org/doc/install

Then, once that is working, all you need to do is: 

```
$ go get github.com/gabeguz/gobot/cmd/gobot
```

Running a bot
-------------

Currently, the command I use is: 

```
$ $GOPATH/bin/gobot -host="jabber.my.domain:5222" -user="gabe@jabber.my.domain" -pass="hunter2" -room="#gobottest@conference.jabber.my.domain" -name="gobot"
```

Creating A Plugin
-----------------

Creating a plugin is pretty easy, you just need to import the xmppbot library: 

```
import (
	"github.com/gabeguz/gobot/bot"
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

One last step, edit gobot/cmd/gobot.go and add a call to your new plugin
(importing if necessary) to the []bot.Plugin slice:

```
[]bot.Plugin{
	myplugin.MyPlugin{},
}
```

That's it!
