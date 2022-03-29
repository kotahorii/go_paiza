package question

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvenOrOdd(t *testing.T) {
	x1 := 3
	x2 := 4
	x3 := -3
	x4 := -4
	x5 := 0

	var ideal, actual string

	actual = EvenOrOdd(x1)
	ideal = "x is positive odd number"
	require.Equal(t, actual, ideal)

	actual = EvenOrOdd(x2)
	ideal = "x is positive even number"
	require.Equal(t, actual, ideal)

	actual = EvenOrOdd(x3)
	ideal = "x is negative odd number"
	require.Equal(t, actual, ideal)

	actual = EvenOrOdd(x4)
	ideal = "x is negative even number"
	require.Equal(t, actual, ideal)

	actual = EvenOrOdd(x5)
	ideal = "x is positive even number"
	require.Equal(t, actual, ideal)
}

func TestReturnGrade(t *testing.T) {
	var result, ideal string
	score1 := 80
	score2 := 70
	score3 := 60
	score4 := 50

	// v1 test
	result = ReturnGradeV1(score1)
	ideal = "pass"
	require.Equal(t, result, ideal)

	result = ReturnGradeV1(score4)
	ideal = "not pass"
	require.Equal(t, result, ideal)

	// v2 test
	result = ReturnGradeV2(score1)
	ideal = "awesome"
	require.Equal(t, result, ideal)

	result = ReturnGradeV2(score2)
	ideal = "great"
	require.Equal(t, result, ideal)

	result = ReturnGradeV2(score4)
	ideal = "bad"
	require.Equal(t, result, ideal)

	// v3 test
	result = ReturnGradeV3(score1)
	ideal = "great"
	require.Equal(t, result, ideal)

	result = ReturnGradeV3(score2)
	ideal = "good"
	require.Equal(t, result, ideal)

	result = ReturnGradeV3(score3)
	ideal = "not bad"
	require.Equal(t, result, ideal)

	result = ReturnGradeV3(score4)
	ideal = "bad"
	require.Equal(t, result, ideal)
}

func TestMidAndFinalExamGrade(t *testing.T) {
	var (
		mid, final    int
		ideal, result string
	)

	mid, final = 70, 70
	result = MidAndFinalExamGrade(mid, final)
	ideal = "pass"
	require.Equal(t, result, ideal)

	mid, final = 50, 80
	result = MidAndFinalExamGrade(mid, final)
	ideal = "pass"
	require.Equal(t, result, ideal)

	mid, final = 10, 90
	result = MidAndFinalExamGrade(mid, final)
	ideal = "pass"
	require.Equal(t, result, ideal)

	mid, final = 40, 40
	result = MidAndFinalExamGrade(mid, final)
	ideal = "not pass"
	require.Equal(t, result, ideal)
}

func TestIsOpenHospital(t *testing.T) {
	var (
		dayOfWeek, time int
		result, ideal   bool
	)

	dayOfWeek = 0
	time = 0
	result = IsOpenHospital(dayOfWeek, time)
	ideal = false
	require.Equal(t, result, ideal)

	dayOfWeek = 1
	time = 0
	result = IsOpenHospital(dayOfWeek, time)
	ideal = true
	require.Equal(t, result, ideal)

	dayOfWeek = 2
	time = 0
	result = IsOpenHospital(dayOfWeek, time)
	ideal = false
	require.Equal(t, result, ideal)
}

func TestXAndYCondition(t *testing.T) {
	var (
		x, y             int
		actual, expected string
	)

	x, y = 2, 4
	actual = XAndYCondition(x, y)
	expected = "x is smaller than y and x and y are both even number."
	require.Equal(t, actual, expected)

	x, y = -5, -5
	actual = XAndYCondition(x, y)
	expected = "x equals to y and x and y are both negative value"
	require.Equal(t, actual, expected)

	x, y = 101, 35
	actual = XAndYCondition(x, y)
	expected = "x is less than or equals to 10\n" +
		"or x is greater than or equals to 100\n" +
		"and y is greater than or equals to 10\n" +
		"and y is less than or equals to 100"

	require.Equal(t, actual, expected)

	x, y = 2, -5
	actual = XAndYCondition(x, y)
	expected = "x and y are not x is negative and y is negative"
	require.Equal(t, actual, expected)
}

func TestNumberOfDays(t *testing.T) {
	var (
		month            int
		actual, expected string
	)

	month = 4
	actual = NumberOfDays(month)
	expected = "30"
	require.Equal(t,actual, expected)

	month = 5
	actual = NumberOfDays(month)
	expected = "31"
	require.Equal(t,actual, expected)

	month = 2
	actual = NumberOfDays(month)
	expected = "28"
	require.Equal(t,actual, expected)

	month = 13
	actual = NumberOfDays(month)
	expected = "It's not month"
	require.Equal(t,actual, expected)
}
