package handler

import (
	"errors"
)

func InputHandler(args []string) error {
	if len(args) < 2 {
		errors.New("not enough arguments")
	}
	return nil
}
