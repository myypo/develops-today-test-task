package verbose

type VerboseError interface {
	error
	Verbose() error
}

type verboseError struct {
	techErr error
	userErr error
}

func NewVerboseError(techErr, userErr error) VerboseError {
	return &verboseError{
		techErr: techErr,
		userErr: userErr,
	}
}

func (e *verboseError) Error() string {
	return e.userErr.Error()
}

func (e *verboseError) Verbose() error {
	return e.techErr
}
