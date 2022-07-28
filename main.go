package main

import (
	"errors"
	"fmt"
)

type Implementation struct {
	StuffToProcess string
}

func (i *Implementation) ProcessStuff() (string, error) {
	if i.StuffToProcess == "" {
		return "", errors.New("stuff was empty")
	}

	return fmt.Sprintf("I have processed the stuff, it was %s", i.StuffToProcess), nil
}

func main() {}
