package selftaught

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarysearch(t *testing.T) {
	a_list := []int{1, 2, 3, 4, 5, 6}
	target := 2
	actual := BinarySearch(a_list, target)
	expected := true

	require.Equal(t, expected, actual)
}
