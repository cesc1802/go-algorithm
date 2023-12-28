package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name           string
		num            int
		expectedResult bool
	}{
		{
			name:           "a happy number",
			num:            1,
			expectedResult: true,
		},
		{
			name:           "is not a happy number",
			num:            5,
			expectedResult: false,
		},
		{
			name:           "a happy number",
			num:            19,
			expectedResult: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isHappy(tc.num)
			assert.Equal(t, tc.expectedResult, got)
		})
	}

}
