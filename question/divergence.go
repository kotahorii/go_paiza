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

func XAndYCondition(x, y int) string {
	if x < y && (x%2 == 0 && y%2 == 0) {
		return "x is smaller than y and x and y are both even number."
	}
	if x == y && (x < 0 && y < 0) {
		return "x equals to y and x and y are both negative value"
	}
	if (x <= 10 || x >= 100) && y >= 10 && y <= 100 {
		return "x is less than or equals to 10\n" +
			"or x is greater than or equals to 100\n" +
			"and y is greater than or equals to 10\n" +
			"and y is less than or equals to 100"

	}
	if !(x < 0 && y < 0) {
		return "x and y are not x is negative and y is negative"
	}
	return ""
}

func DisplayHoliday(month int) {
	switch month {
	case 1:
		fmt.Println("ganjitsu")
		fmt.Println("seijinn no hi")
		fallthrough
	case 2:
		fmt.Println("kenkoku kinnenbi")
		fallthrough
	case 3:
		fmt.Println("shuubunn no hi")
		fallthrough
	case 4:
		fmt.Println("showa no hi")
		fallthrough
	case 5:
		fmt.Println("kenpou kinennbi")
		fmt.Println("midori no hi")
		fmt.Println("kodomo no hi")
		fallthrough
	case 7:
		fmt.Println("umi no hi")
		fallthrough
	case 9:
		fmt.Println("keiro no hi")
		fmt.Println("shuubunn no hi")
		fallthrough
	case 10:
		fmt.Println("taiiku no hi")
		fallthrough
	case 11:
		fmt.Println("bunnka no hi")
		fmt.Println("kinnro kansha no hi")
		fallthrough
	case 12:
		fmt.Println("tennno tanjobi")
	}
}

func NumberOfDays(month int) string {
	switch month {
	case 4, 6, 9, 11:
		return "30"
	case 1, 3, 5, 7, 8, 10, 12:
		return "31"
	case 2:
		return "28"
	default:
		return "It's not month"
	}
}
