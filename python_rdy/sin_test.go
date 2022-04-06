package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSin(t *testing.T) {
	actual := CalcSin(180)
	expected := 0.0
	require.Equal(t, expected, actual)

	actual = CalcSin(30)
	expected = 0.5
	require.Equal(t, expected, actual)

	actual = CalcSin(60)
	expected = 0.9
	require.Equal(t, expected, actual)

	actual = CalcSin(90)
	expected = 1.0
	require.Equal(t, expected, actual)
}
