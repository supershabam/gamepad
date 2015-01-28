package main

import (
	"github.com/supershabam/gamepad"
	"log"
	"time"
)

func main() {
	in := make(chan gamepad.Event)
	g := gamepad.NewGamepad(in)

	next := make(chan gamepad.Event, 0)
	g.Notify(next, gamepad.Left, gamepad.Right, gamepad.Up)

	in <- gamepad.Event{Button: gamepad.Up, Pressed: true}
	in <- gamepad.Event{Button: gamepad.Up, Pressed: false}

	select {
	case e := <-next:
		log.Printf("%v", e)
		g.Stop(next)
	case <-time.After(time.Second):
		log.Printf("%+v", g.State())
	}

	select {
	case e := <-next:
		log.Printf("%v", e)
		g.Stop(next)
	case <-time.After(time.Second):
		log.Printf("%+v", g.State())
	}
}
