package question

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisplaySPAM10Times(t *testing.T) {
	var actual, expected []string
	actual = DisplaySPAM10Times()
	expected = []string{
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
		"SPAM",
	}
	require.Equal(t, actual, expected)
}

func TestReturnFactorial(t *testing.T) {
	var actual, expected int
	actual = ReturnFactorial(7)
	expected = 5040

	require.Equal(t, actual, expected)
}
