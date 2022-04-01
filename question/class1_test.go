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

func TestGetAllCount(t *testing.T) {
	coincase := NewCoinCase()
	coincase.AddCoins(1, 3)
	coincase.AddCoins(5, 2)
	coincase.AddCoins(10, 3)
	coincase.AddCoins(50, 1)
	coincase.AddCoins(100, 2)
	coincase.AddCoins(500, 3)

	actual := coincase.GetAllCount()
	expected := 14
	require.Equal(t, actual, expected)
}

func TestGetAmount(t *testing.T) {
	coincase := NewCoinCase()
	coincase.AddCoins(1, 3)
	coincase.AddCoins(5, 2)
	coincase.AddCoins(10, 3)
	coincase.AddCoins(50, 1)
	coincase.AddCoins(100, 2)
	coincase.AddCoins(500, 3)

	actual := coincase.GetAmount()
	expected := 1793
	require.Equal(t, actual, expected)
}

func TestGetSomeAmount(t *testing.T) {
	coincase := NewCoinCase()
	coincase.AddCoins(500, 3)

	actual, _ := coincase.GetSomeAmount(500)
	expected := 1500
	require.Equal(t, actual, expected)
}
