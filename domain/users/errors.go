package users

type ErrUserInfoAlreadyRegistered struct {
}

func (err ErrUserInfoAlreadyRegistered) Error() string {
	return "User with this information  already registered"
}
