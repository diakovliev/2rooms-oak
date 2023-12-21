package utils_test

import (
	"errors"
	"testing"

	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/stretchr/testify/assert"
)

func TestSplitToLines(t *testing.T) {

	type testCase struct {
		input    string
		maxLen   int
		expected []string
		err      error
	}

	testCases := []testCase{
		{
			input:  "this is test string",
			maxLen: 5,
			err:    utils.ErrUnableToSplit,
		},
		{
			input:  "this is test string",
			maxLen: 10,
			expected: []string{
				"this is",
				"test",
				"string",
			},
		},
		{
			input:  "this is test string",
			maxLen: 20,
			expected: []string{
				"this is test string",
			},
		},
		{
			input:  " This is a test string",
			maxLen: 19,
			expected: []string{
				" This is a test",
				"string",
			},
		},
		{
			input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			maxLen: 10,
			err:    utils.ErrUnableToSplit,
		},
		{
			input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			maxLen: 20,
			expected: []string{"Lorem ipsum dolor",
				"sit amet,",
				"consectetur",
				"adipiscing elit.",
			},
		},
		{
			input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			maxLen: 30,
			expected: []string{
				"Lorem ipsum dolor sit amet,",
				"consectetur adipiscing elit.",
			},
		},
		{
			input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			maxLen: 40,
			expected: []string{
				"Lorem ipsum dolor sit amet, consectetur",
				"adipiscing elit.",
			},
		},
		{
			input:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			maxLen: 50,
			expected: []string{
				"Lorem ipsum dolor sit amet, consectetur adipiscing",
				"elit.",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := utils.SplitToLines(tc.input, tc.maxLen)
			if tc.err != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tc.err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}
