package question

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCounts(t *testing.T) {
	coincase := NewCoinCase()
	if err := coincase.AddCoins(1, 3); err != nil {
		log.Fatal(err)
	}

	actual, _ := coincase.GetCounts(1)
	expected := 3
	require.Equal(t, actual, expected)

	actual, _ = coincase.GetCounts(10)
	expected = 0
	require.Equal(t, actual, expected)

	_, err := coincase.GetCounts(30)
	require.Error(t, err)
}
