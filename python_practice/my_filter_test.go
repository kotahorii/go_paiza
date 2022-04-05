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
