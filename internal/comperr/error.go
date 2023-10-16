package comperr

import "nur/internal/position"

type Error struct {
	Message string
	position.Position
}

func (e Error) Error() string {
	return e.Message
}
