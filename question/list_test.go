package question

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwoTimes(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	actual := TwoTimes(numbers)
	expected := []int{2, 4, 6, 8, 10}

	require.Equal(t, actual, expected)
}

func TestReverseNums(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	actual := ReverseNums(numbers)
	expected := []int{5, 4, 3, 2, 1}

	require.Equal(t, actual, expected)
}

func TestDivideEvenOdd(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	actual1, actual2 := DivideEvenOdd(numbers)

	require.Equal(t, actual1, []int{2, 4})
	require.Equal(t, actual2, []int{1, 3, 5})
}

func TestCalcProduct(t *testing.T) {
	num1, num2 := 5, 4
	actual := CalcProduct(num1, num2)
	expected := 20

	require.Equal(t, actual, expected)
}
