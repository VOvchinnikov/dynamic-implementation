package main

import (
	"errors"
	"fmt"
)

type Interface interface { // re-declare it here for the getter method
	ProcessStuff() (string, error)
}

type Implementation struct {
	StuffToProcess string
}

func NewImplementation() Interface {
	return &Implementation{}
}

func (i *Implementation) ProcessStuff() (string, error) {
	if i.StuffToProcess == "" {
		return "", errors.New("stuff was empty")
	}

	return fmt.Sprintf("I have processed the stuff, it was %s", i.StuffToProcess), nil
}

func main() {}
