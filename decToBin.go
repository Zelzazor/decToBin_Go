package main

import (
	"errors"
	"strconv"
)

func decToBin(number int) (string, error) {
	if number < 0 {
		return "0", errors.New("number is out of range")
	} else if number == 0 {
		return "0", nil
	} else {
		var result = ""
		var char = ""
		for number > 0 {
			char = strconv.Itoa(number % 2)
			result = char + result
			number /= 2
		}

		return result, nil
	}
}
