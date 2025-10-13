package main

import (
	"fmt"
)

func (u User) Label() string {
	return fmt.Sprintf("%s (%d)", u.Name, u.Age)
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}
