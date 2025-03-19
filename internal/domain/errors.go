package domain

type Error struct {
	text string
}

func (e *Error) Error() string { return e.text }

var (
	ErrEmptyName             = &Error{"empty name"}
	ErrAlreadyBlocked        = &Error{"already blocked"}
	ErrAlreadyUnblocked      = &Error{"already unblocked"}
	ErrAccountBlocked        = &Error{"account is blocked"}
	ErrAlreadyApplied        = &Error{"already appplied"}
	ErrNotEnoughMoney        = &Error{"not enough money. Work harder bro"}
	ErrCannotResolveCategory = &Error{"cannot resolve category"}
)
