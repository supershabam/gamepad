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
			b := e.Button()
			switch b {
			case DPadUp:
				g.s.Up = e.Bool()
			case DPadDown:
				g.s.Down = e.Bool()
			case DPadLeft:
				g.s.Left = e.Bool()
			case DPadRight:
				g.s.Right = e.Bool()
			}
			for n, btns := range g.nm {
				for _, btn := range btns {
					if b == btn {
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
