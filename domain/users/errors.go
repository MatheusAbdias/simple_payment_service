package users

type ErrEmailAlreadyRegistered struct {
}

func (err ErrEmailAlreadyRegistered) Error() string {
	return "User with this email already registered"
}

type ErrDocumentAlreadyRegistered struct {
}

func (err ErrDocumentAlreadyRegistered) Error() string {
	return "User with this document already registered"
}
