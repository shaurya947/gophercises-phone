package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type normalizeTestCase struct {
	in, out   string
	shouldErr bool
}

func TestNormalizePhone(t *testing.T) {
	testCases := []normalizeTestCase{
		{in: "", out: "", shouldErr: true},
		{in: "  4", out: "", shouldErr: true},
		{in: "  4", out: "", shouldErr: true},
		{in: "1234567890", out: "1234567890", shouldErr: false},
		{in: "(123)456-7892", out: "1234567892", shouldErr: false},
		{in: "(  123) 456-78 93", out: "1234567893", shouldErr: false},
		{in: "4(123)456-7892", out: "", shouldErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			normalized, err := normalizePhone(tc.in)
			require.Equal(t, tc.out, normalized)
			if tc.shouldErr {
				require.NotNil(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
