package pythonpractice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEven(t *testing.T) {
	num := 3
	actual := IsEven(num)
	expected := false
	require.Equal(t, actual, expected)
}

func TestMyFilter(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	f := IsEven
	actual := MyFilter(f, list)
	expected := []int{2, 4}
	require.Equal(t, actual, expected)
}

func TestPartition(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	f := IsEven

	actual := Partition(f, list)
	expected := [][]int{{2, 4}, {1, 3, 5}}
	require.Equal(t, actual, expected)

}

func TestPartitionScore(t *testing.T) {
	scoreList := map[string]int{
		"Beth":  43,
		"Kathy": 80,
		"Mark":  56,
		"Mary":  70,
		"Susie": 68,
	}
	f := IsScoreGreaterThan60

	actual := PartitionScore(f, scoreList)
	expected := []map[string]int{
		{
			"Kathy": 80,
			"Mary":  70,
			"Susie": 68,
		},
		{
			"Beth": 43,
			"Mark": 56,
		},
	}
	require.Equal(t, actual, expected)
}

func TestGrange(t *testing.T) {
	actual := Grange(0, 10, 1)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	require.Equal(t, actual, expected)

	actual = Grange(1, 5, 1)
	expected = []int{1, 2, 3, 4}
	require.Equal(t, actual, expected)

	actual = Grange(0, 10, 2)
	expected = []int{0, 2, 4, 6, 8}
	require.Equal(t, actual, expected)

	actual = Grange(0, -5, -1)
	expected = []int{0, -1, -2, -3, -4}
	require.Equal(t, actual, expected)

	actual = Grange(0, 0, 1)
	expected = []int{}
	require.Equal(t, actual, expected)

	actual = Grange(1, 0, 1)
	expected = []int{}
	require.Equal(t, actual, expected)

	actual = Grange(0, 1, -1)
	expected = []int{}
	require.Equal(t, actual, expected)
}
