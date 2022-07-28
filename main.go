package main

import (
	"errors"
	"fmt"
)

type Interface interface { // re-declare it here for the getter method
	ProcessStuff(string) (string, error)
}

type Implementation struct{}

func NewImplementation() Interface {
	return &Implementation{}
}

func (i *Implementation) ProcessStuff(stuff string) (string, error) {
	if stuff == "" {
		return "", errors.New("stuff was empty")
	}

	return fmt.Sprintf("I have processed the stuff, it was %s", stuff), nil
}

func main() {}
