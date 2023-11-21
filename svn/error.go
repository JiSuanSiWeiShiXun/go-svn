package svn

import "errors"

var (
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrInvalidURL           = errors.New("invalid URL")
	ErrInvalidSVNPath       = errors.New("invalid svn path")
	ErrRepositoryNotFound   = errors.New("repository not found")
	ErrUnknown              = errors.New("unknown error")
)
