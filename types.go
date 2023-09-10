package main

import "math/rand"

type Account struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	BankNumber int64  `json:"bankNumber"`
	Balance    int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:         rand.Intn(10000),
		FirstName:  firstName,
		LastName:   lastName,
		BankNumber: int64(rand.Intn(10000000)),
	}
}
