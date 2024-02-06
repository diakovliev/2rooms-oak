package button

type State string

const (
	Disabled State = "disabled"
	Down     State = "down"
	Up       State = "up"
)

type Focus string

const (
	Focused    Focus = "focused"
	NonFocused Focus = "nonfocused"
)
