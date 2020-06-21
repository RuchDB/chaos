package redis

import (
	"testing"
)

func TestFSM(t *testing.T) {
	var name = "SET hello 1 "
	t.Run(name, func(t *testing.T) {
		var b = []byte(name)
		var redisFsm = Create(' ')
		redisFsm.Parse(b)
		var action = redisFsm.Action()
		var args = redisFsm.Args()
		if action != "SET" {
			t.Errorf("ACTION: Expected SET, but found %s", action)
		}
		if args[0] != "hello" {
			t.Errorf("ARGS[0]: Expected hello, but found %s", args[0])
		}
		if args[1] != "1" {
			t.Errorf("ARGS[1]: Expected hello, but found %s", args[1])
		}
	})
	name = "gET world 12 34 56 78"
	t.Run(name, func(t *testing.T) {
		var b = []byte(name)
		var redisFsm = Create(' ')
		redisFsm.Parse(b)
		var action = redisFsm.Action()
		var args = redisFsm.Args()
		if action != "GET" {
			t.Errorf("ACTION: Expected SET, but found %s", action)
		}
		if len(args) != 5 {
			t.Errorf("ARGS: Expected length 5, but found %d", len(args))
		}
	})
}
