package P0412_FizzBuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "Fizz"}, fizzBuzz(3))
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, fizzBuzz(5))
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, fizzBuzz(15))
}
