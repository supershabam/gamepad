package gamepad

type Button int64

const (
	DPadUp Button = iota
	DPadDown
	DPadLeft
	DPadRight
)

type Event interface {
	Bool() bool
	Button() Button
}

type UpDPadEvent struct{ Value bool }

func (e UpDPadEvent) Bool() bool     { return e.Value }
func (e UpDPadEvent) Button() Button { return DPadUp }

type DownDPadEvent struct{ Value bool }

func (e DownDPadEvent) Bool() bool     { return e.Value }
func (e DownDPadEvent) Button() Button { return DPadDown }

type LeftDPadEvent struct{ Value bool }

func (e LeftDPadEvent) Bool() bool     { return e.Value }
func (e LeftDPadEvent) Button() Button { return DPadLeft }

type RightDPadEvent struct{ Value bool }

func (e RightDPadEvent) Bool() bool     { return e.Value }
func (e RightDPadEvent) Button() Button { return DPadRight }
