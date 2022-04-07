package pythonrdy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganizeList(t *testing.T) {
	l := [][]string{
		{"Py", "PyA"},
		{"Py", "PyB"},
		{"Py", "PyC"},
		{"Pe", "PeA"},
		{"Pe", "PeB"},
		{"Pe", "PeC"},
	}
	actual := OrganizeList(l)
	expected := map[string][]string{
		"Py": {"PyA", "PyB", "PyC"},
		"Pe": {"PeA", "PeB", "PeC"},
	}
	require.Equal(t, expected, actual)
}
