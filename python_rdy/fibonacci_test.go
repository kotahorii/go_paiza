package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFib(t *testing.T) {
	actual := Fib(10)
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	require.Equal(t, actual, expected)
}
