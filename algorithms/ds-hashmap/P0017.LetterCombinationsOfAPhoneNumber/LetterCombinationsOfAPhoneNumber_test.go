package p0017lettercombinationsofaphonenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	letterCombinations("345")
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal(t, []string{}, letterCombinations(""))
}
