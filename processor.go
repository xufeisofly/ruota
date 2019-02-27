package ruota

type RProcessor interface {
	Call(RTransport, RSerializer) error
}
