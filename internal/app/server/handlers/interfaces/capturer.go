package interfaces

type Capturer interface {
	Capture([]byte) error
}
