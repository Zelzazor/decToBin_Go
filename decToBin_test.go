package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultZero(t *testing.T) {

	shouldBe0, err := decToBin(0)
	assert.Equal(t, "0", shouldBe0, "0 in binary should be 0")
	assert.Equal(t, nil, err, "Error should be nil")

}

func TestResultOne(t *testing.T) {
	shouldBe1, err := decToBin(1)
	assert.Equal(t, "1", shouldBe1, "1 in binary should be 1")
	assert.Equal(t, nil, err, "Error should be nil")
}

func TestResultEight(t *testing.T) {
	shouldBe8, err := decToBin(8)
	assert.Equal(t, "10", shouldBe8, "2 in binary should be 10")
	assert.Equal(t, nil, err, "Error should be nil")
}

func TestResult1024(t *testing.T) {
	shouldBe1024, err := decToBin(1024)
	assert.Equal(t, "10000000000", shouldBe1024, "1024 in binary should be 10000000000")
	assert.Equal(t, nil, err, "Error should be nil")
}

func TestResult6278956(t *testing.T) {
	shouldBe6278956, err := decToBin(6278956)
	assert.Equal(t, "010111111100111100101100", shouldBe6278956, "6278956 in binary should be 010111111100111100101100")
	assert.Equal(t, nil, err, "Error should be nil")
}

func TestResult2147483647(t *testing.T) {
	shouldBe2147483647, err := decToBin(2147483647)
	assert.Equal(t, "1111111111111111111111111111111", shouldBe2147483647, "2147483647 in binary should be 1111111111111111111111111111111")
	assert.Equal(t, nil, err, "Error should be nil")
}

func TestShouldNotBeNegativeOne(t *testing.T) {
	result, err := decToBin(-1)
	if assert.Error(t, err) {
		assert.Equal(t, "0", result)
		assert.Equal(t, errors.New("number is out of range"), err, "Error is not the same")
	}
}

func TestShouldNotBeNegativeSixtyFour(t *testing.T) {
	result, err := decToBin(-64)
	if assert.Error(t, err) {
		assert.Equal(t, "0", result)
		assert.Equal(t, errors.New("number is out of range"), err, "Error is not the same")
	}
}
