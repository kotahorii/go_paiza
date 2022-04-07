package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvTemperature(t *testing.T) {
	actual := ConvTemperature("5c")
	expected := "41.0F"
	require.Equal(t, expected, actual)

	actual = ConvTemperature("41f")
	expected = "5.0C"
	require.Equal(t, expected, actual)
}
