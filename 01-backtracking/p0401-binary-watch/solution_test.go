package solution

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase struct {
	input    int
	expected []string
}

func TestReadBinaryWatch(t *testing.T) {
	testcases := []testcase{
		{
			input:    1,
			expected: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			input:    9,
			expected: []string{},
		},
	}

	for _, testcase := range testcases {
		sort.Strings(testcase.expected)
		actual := readBinaryWatch(testcase.input)
		sort.Strings(actual)
		require.Equal(t, testcase.expected, actual)
	}
}
