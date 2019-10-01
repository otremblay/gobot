// Copyright 2014 Gabriel Guzman <gabe@lifewaza.com>
// All rights reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package dm gives the bot the ability to recognize messages
//  sent to himself, and reply accordingly
package dm

import (
	"strings"

	"github.com/gabeguz/gobot"
)

// Helper struct that will implement the helper interface
type DirectMessage struct {
}

func (p DirectMessage) Name() string {
	return "Direct Message v1.0"
}

// Send allows the bot to send a message to this helper
func (p DirectMessage) Execute(message gobot.Message, cb gobot.Bot) error {
	reply := dm(message.Body(), cb.Name())
	if len(reply) > 0 {
		cb.Reply(message, reply)
	}
	return nil
}

func dm(message, nick string) string {
	reply := ""
	if strings.Contains(strings.ToLower(message), strings.ToLower(nick)) {
		reply = ("my ears are burning")
	}
	return reply
}
