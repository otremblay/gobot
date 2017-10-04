// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package chatlog logs all chat messages to a text file
package chatlog

import (
	"fmt"
	"log"
	"os"

	"github.com/gabeguz/gobot"
	gb "github.com/gabeguz/gobot/bots/gobot"
	sb "github.com/gabeguz/gobot/bots/slack"
	xmpp "github.com/gabeguz/gobot/bots/xmpp"
)

// Helper struct that will implement the Helper interface
type ChatLog struct {
	Filename string
}

func (c ChatLog) Name() string {
	return "Chatlog v1.0"
}

var users = map[string]string{}

// Send allows the bot to send a message to this helper
func (c ChatLog) Execute(message gobot.Message, bot gobot.Bot) error {
	if c.Filename == "" {
		c.Filename = "/tmp/chatlog"
	}
	b2 := bot.(gb.Gobot)
	switch b3 := b2.InternalBot().(type) {
	case *sb.Bot:
		tmpmsg := message.(sb.Message)
		mess := tmpmsg.EffectiveMessage()
		if mess.Channel != b3.Opt.Room {
			return nil
		}
		var n string
		if user, ok := users[mess.User]; ok {
			n = user
		} else {
			u, err := b3.Client().GetUserInfo(mess.User)
			if err == nil {
				n = u.Name
				users[mess.User] = n
			}
		}
		c.Logit(n, message.Body())
		for _, a := range mess.Attachments {
			c.Logit(message.From(), a.Title)
			c.Logit(message.From(), a.Text)
		}
	case *xmpp.Bot:
		if message.From() != b3.Opt.Room {
			return nil
		}
		c.Logit(message.From(), message.Body())
	default:
		c.Logit(message.From(), message.Body())
	}
	return nil
}

func (c ChatLog) Logit(author, message string) {
	f, err := os.OpenFile(c.Filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	if message == "" {
		return
	}
	log.SetOutput(f)
	log.Println(fmt.Sprintf("%s: %s", author, message))
}
