package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizePhone(t *testing.T) {
	normalized, err := normalizePhone("")
	require.Empty(t, normalized)
	require.NotNil(t, err)

	normalized, err = normalizePhone("  4")
	require.Empty(t, normalized)
	require.NotNil(t, err)

	normalized, err = normalizePhone("1234567890")
	require.Equal(t, "1234567890", normalized)
	require.Nil(t, err)

	normalized, err = normalizePhone("(123)456-7892")
	require.Equal(t, "1234567892", normalized)
	require.Nil(t, err)

	normalized, err = normalizePhone("(  123) 456-78 93")
	require.Equal(t, "1234567893", normalized)
	require.Nil(t, err)

	normalized, err = normalizePhone("4(123)456-7892")
	require.Empty(t, normalized)
	require.NotNil(t, err)
}
