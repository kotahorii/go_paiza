package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSortNumInTextFile(t *testing.T) {
	actual := SortNumInTextFile("../text.txt")
	expected := []int{732, 542, 341, 64}
	require.Equal(t, expected, actual)
}
