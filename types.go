package main

import (
	"math/rand"
)

type User struct {
	PK          int    `json:"pk"`
	FirstName   string `json:"fistName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func NewUser(firstName, lastName, phoneNumber, password string) *User {
	return &User{
		PK:          rand.Intn(10000),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}
