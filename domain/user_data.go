package domain

import (
	"context"
)

// var (
// 	ErrAccountNotFound = errors.New("account not found")

// 	ErrAccountOriginNotFound = errors.New("account origin not found")

// 	ErrAccountDestinationNotFound = errors.New("account destination not found")

// 	ErrInsufficientBalance = errors.New("origin account does not have sufficient balance")
// )

type ID int

func (u ID) Int() int {
	return int(u)
}

type (
	UserDataRepository interface {
		Create(context.Context, UserData) (UserData, error)
	}

	UserData struct {
		id ID
	}
)

func NewUserData(ID ID) UserData {
	return UserData{
		id: ID,
	}
}

func (u UserData) ID() ID {
	return u.id
}
