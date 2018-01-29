package xmpp

import (
	"testing"
)

func TestNew(t *testing.T) {
	bot := New("localhost:12345", "bob", "hunter2", "chatroom", "robot")
	expected := "localhost:12345"
	if bot.Opt.Host != expected {
		t.Errorf("Invalid host %s, expected %s", bot.Opt.Host, expected)
	}
	expected = "xmpp"
	if bot.Opt.Protocol != expected {
		t.Errorf("Invalid protocol %s, expected %s", bot.Opt.Protocol, expected)
	}
	expected = "bob"
	if bot.Opt.User != expected {
		t.Errorf("Invalid user %s, expected %s", bot.Opt.User, expected)
	}
	expected = "hunter2"
	if bot.Opt.Password != expected {
		t.Errorf("Invalid password %s, expected %s", bot.Opt.Password, expected)
	}
	expected = "chatroom"
	if bot.Opt.Room != expected {
		t.Errorf("Invalid protocol %s, expected %s", bot.Opt.Room, expected)
	}
	expected = "robot"
	if bot.Opt.Resource != expected {
		t.Errorf("Invalid protocol %s, expected %s", bot.Opt.Resource, expected)
	}
}
