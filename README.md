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
./seedbot -nick="mybot" -notls=true -password="hunter2" -room="#seedbottest@conference.desktop.my.domain" -server="desktop.my.domain:5222" -session=true -username="gabe@desktop.my.domain"
```

Creating A Plugin
-----------------

Creating a plugin is pretty easy, you just need to...

TODO -- Fill out this section!

