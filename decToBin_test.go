package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	shouldBe0, err := decToBin(0)
	assert.Equal(t, "0", shouldBe0, "0 in binary")
	shouldBe1, err := decToBin(1)
	assert.Equal(t, "1", shouldBe1, "1 in binary should be 1")
	shouldBe8, err := decToBin(8)
	assert.Equal(t, "10", shouldBe8, "2 in binary should be 10")
	shouldBe1024, err := decToBin(1024)
	assert.Equal(t, "10000000000", shouldBe1024, "1024 in binary should be 10000000000")
	shouldBe6278956, err := decToBin(6278956)
	assert.Equal(t, "010111111100111100101100", shouldBe6278956, "6278956 in binary should be 010111111100111100101100")
	shouldBe2147483647, err := decToBin(2147483647)
	assert.Equal(t, "1111111111111111111111111111111", shouldBe2147483647, "2147483647 in binary should be 1111111111111111111111111111111")
	result, err := decToBin(-1)
	if assert.Error(t, err) {
		assert.Equal(t, "0", result)
		assert.Equal(t, errors.New("number is out of range"), err, "Error is not the same")
	}
	result2, err := decToBin(-64)
	if assert.Error(t, err) {
		assert.Equal(t, "0", result2)
		assert.Equal(t, errors.New("number is out of range"), err, "Error is not the same")
	}

}
