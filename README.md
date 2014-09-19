seedbot
=======

seedbot is a Simple XMPP bot written in golang  

Initail work based on https://github.com/mattn/go-xmpp/blob/master/example/example.go

seedbot is under active development, the api may change w/out warning.  

By itself, seedbot just sits and listens for incoming XMPP messages.  When a message is received, seedbot passes the message on to any number of plugins.  The plugins provide all the features of the chatbot.

Currently, the following plugins exist: 

1. dm - seedbot recognizes his nick and replies that his ears are burning when he is mentioned. 

2. quote - seedbot understands the quote command.  Currently 3 types of quotes are supported: 
	code quote - quotes about programming
	integration quote - quotes about front end integration
	admin quote - quotes about systems administration

3. chatlog - seedbot logs all messages to a text file. Currently '/tmp/chatlog'

4. stathat - seedbot counts each message, and logs it as a single 'hit' at stathat.

5. keyword - seedbot scans a message for work counts

6. echo - seedbot echoes a message back to the room

Running a bot
-------------

Currently, the command I use is: 

```
./seedbot -host="jabber.my.domain:5222" -user="gabe@jabber.my.domain" -pass="hunter2" -room="#seedbottest@conference.jabber.my.domain" -name="seedbot"
```

Creating A Plugin
-----------------

Creating a plugin is pretty easy, you just need to import the xmppbot library: 

```
import (
	"github.com/seedboxtech/xmppbot"
)
```

And then implement 2 exported methods so that you conform to the seedbotplugin interface: 

```
type MyPlugin struct {}

func (p MyPlugin) Name() string {
	return "MyPlugin v1.0"
}

func (p MyPlugin) Execute(msg xmppbot.Message, bot xmppbot.Bot) error {
	if msg.From() != bot.FullName() {
		bot.Send("Why, hello!")
	}
	return nil
}
```

One last step, edit seedbot/seedbot.go and add a call to your new plugin (importing if necessary) to the []seedbotplugin.Plugin slice:

```
[]seedbotplugin.Plugin{
	seedbotplugin.MyPlugin{},
}
```

That's it!
