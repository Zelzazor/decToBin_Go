package main

import (
	"errors"
)

func decToBin(number int) (string, error) {
	if number < 0 {
		return "0", errors.New("number is out of range")
	}

	return "0", nil
}
