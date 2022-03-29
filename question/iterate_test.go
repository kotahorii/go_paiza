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
