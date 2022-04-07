package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCombineTwoList(t *testing.T) {
	l1, l2 := []string{"AA", "AB", "AC"}, []string{"BA", "BB", "BC"}
	actual := CombineTwoList(l1, l2)
	expected := [][]string{
		{"AA", "BA"},
		{"AA", "BB"},
		{"AA", "BC"},
		{"AB", "BA"},
		{"AB", "BB"},
		{"AB", "BC"},
		{"AC", "BA"},
		{"AC", "BB"},
		{"AC", "BC"},
	}
	require.Equal(t, expected, actual)
}

func TestNEmptyList(t *testing.T) {
	actual := MakeNEmptyList(3)
	expected := [][]interface{}{{}, {}, {}}
	require.Equal(t, expected, actual)
}
