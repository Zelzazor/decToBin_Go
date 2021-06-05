package main

import (
	"errors"
	"strconv"
)

func reverseString(str string) string {
	strArr := []rune(str)
	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)
}

func decToBin(number int) (string, error) {
	if number < 0 {
		return "0", errors.New("number is out of range")
	} else if number == 0 {
		return "0", nil
	} else {
		var result = ""
		for number > 0 {
			result += strconv.Itoa(number % 2)
			number /= 2
		}
		result = reverseString(result)
		return result, nil
	}
}
