package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGCD(t *testing.T) {
	a, b := 5, 10
	actual := GCD(a, b)
	expected := a
	require.Equal(t, actual, expected)

	// LCM test
	actual = a * b / GCD(a, b)
	expected = 10
	require.Equal(t, actual, expected)
}
