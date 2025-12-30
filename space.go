package main

type Space int

const (
	Empty Space = iota
	Red
	Yellow
)

func (s Space) Symbol() string {
	switch s {
	case Red:
		return "\033[31m●\033[0m"
	case Yellow:
		return "\033[33m●\033[0m"
	case Empty:
		return "·"
	default:
		return "?"
	}
}

func (s Space) String() string {
	switch s {
	case Red:
		return "\033[31mRed\033[0m"
	case Yellow:
		return "\033[33mYellow\033[0m"
	default:
		return ""
	}
}
