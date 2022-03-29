package question

import "fmt"

func PrintGreater(x, y int) {
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is smaller than or equal to y")
	}
}

func ReturnGreaterValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func EvenOrOdd(x int) string {
	// x is even
	if x%2 == 0 {
		if x >= 0 {
			return "x is positive even number"
		}
		return "x is negative even number"
	}

	// x is odd
	if x > 0 {
		return "x is positive odd number"
	}
	return "x is negative odd number"
}

func ReturnGradeV1(score int) string {
	if score >= 60 {
		return "pass"
	}
	return "not pass"
}

func ReturnGradeV2(score int) string {
	if score >= 80 {
		return "awesome"
	}
	if score >= 60 && score < 80 {
		return "great"
	}
	return "bad"
}

func ReturnGradeV3(score int) string {
	if score >= 80 {
		return "great"
	}
	if score >= 70 && score < 80 {
		return "good"
	}
	if score >= 60 && score < 70 {
		return "not bad"
	}
	return "bad"
}

func MidAndFinalExamGrade(mid, final int) string {
	if (mid >= 60 && final >= 60) ||
		(mid+final >= 130) ||
		(mid+final >= 100 &&
			(mid >= 90 || final >= 90)) {
		return "pass"
	}
	return "not pass"
}

func IsOpenHospital(dayOfWeek, time int) bool {
	if dayOfWeek == 0 {
		return false
	}
	if dayOfWeek == 2 || dayOfWeek == 5 {
		if time == 0 {
			return false
		}
	}
	if dayOfWeek == 3 {
		if time == 2 {
			return false
		}
	}
	if dayOfWeek == 6 {
		if time == 1 || time == 2 {
			return false
		}
	}

	return true
}
