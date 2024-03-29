package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrimeNumber(t *testing.T) {
	actual := PrimeNumber(10)
	expected := []int{2, 3, 5, 7}
	require.Equal(t, expected, actual)
}

func TestNtimes(t *testing.T) {
	actual := NTimes(3)
	expected := 7
	require.Equal(t, expected, actual)

	actual = NTimes(2)
	expected = 6
	require.Equal(t, expected, actual)
}
