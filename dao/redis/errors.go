package redis

import "errors"

var (
	ErrorVoteTimeExpired = errors.New("vote time expired")
	ErrVoteRepeated      = errors.New("vote repeated")
)
