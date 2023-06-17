package tokenprovider

import (
	"errors"
	"food-delivery/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEnCodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ErrEnCodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("error invalid provided"),
		"error invalid provided",
		"ErrInvalidToken",
	)
)
