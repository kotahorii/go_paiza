package question

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisplaySPAM10Times(t *testing.T) {
	actual := DisplaySPAM10Times()
	expected := []string{
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
	actual := ReturnFactorial(7)
	expected := 5040

	require.Equal(t, actual, expected)
}

func TestRepeatAste(t *testing.T) {
	actual := RepeatAste(10)
	expected := "**********"

	require.Equal(t, actual, expected)
}

func TestReturnIterateNum(t *testing.T) {
	actual := ReturnIterateNum(14)
	expected := "01234567890123"

	require.Equal(t, actual, expected)
}
