package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	actual := Sum(10)
	expected := 55
	require.Equal(t, actual, expected)
}
