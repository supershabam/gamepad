package gamepad

type Gamepad struct {
	s  State
	nm map[chan<- Event][]Button
}

func NewGamepad(in <-chan Event) *Gamepad {
	g := &Gamepad{
		s:  State{},
		nm: make(map[chan<- Event][]Button),
	}
	go func() {
		for e := range in {
			switch e.Button {
			case Up:
				g.s.Up = e.Pressed
			case Down:
				g.s.Down = e.Pressed
			case Left:
				g.s.Left = e.Pressed
			case Right:
				g.s.Right = e.Pressed
			}
			for n, btns := range g.nm {
				for _, btn := range btns {
					if e.Button == btn {
						select {
						case n <- e:
						default: // non-blocking
						}
					}
				}
			}
		}
	}()
	return g
}

func (g *Gamepad) Notify(out chan<- Event, btns ...Button) {
	g.nm[out] = btns[:]
}

func (g *Gamepad) Stop(out chan<- Event) {
	delete(g.nm, out)
}

func (g Gamepad) State() State {
	return State{
		Up:    g.s.Up,
		Down:  g.s.Down,
		Left:  g.s.Left,
		Right: g.s.Right,
	}
}
