package interfaces

import "context"

type Capturer interface {
	Capture(context.Context, []byte) error
}
