package question

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoPowers(t *testing.T) {
	num := MyInt(10)
	actual := num.TwoPowers()
	expected := 100

	require.Equal(t, expected, actual)
}

func TestCalcMean(t *testing.T) {
	num := MyInt(10)
	actual := num.CalcMean(MyInt(20))
	expected := MyInt(15)

	require.Equal(t, expected, actual)
}

func TestReturnLargeNum(t *testing.T) {
	num := MyInt(10)
	actual := num.ReturnLargeNum(30)
	expected := 30

	require.Equal(t, actual, expected)
}

func TestReturnLargestNum(t *testing.T) {
	x, y, z := 10, 20, 30
	actual := ReturnLargestNum(x, y, z)
	expected := 30

	require.Equal(t, actual, expected)
}
