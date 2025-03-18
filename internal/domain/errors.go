package domain

type Error struct {
	text string
}

func (e *Error) Error() string { return e.text }

var (
	ErrEmptyName        = &Error{"empty name"}
	ErrAlreadyBlocked   = &Error{"already blocked"}
	ErrAlreadyUnblocked = &Error{"already unblocked"}
)
