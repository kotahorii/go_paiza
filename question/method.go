package question

import "math"

type MyInt int

func (i MyInt) TwoPowers() int {
	return int(math.Pow(float64(i), 2))
}

func (i MyInt) CalcMean(other MyInt) MyInt {
	return (i + other) / 2
}

func (i MyInt) ReturnLargeNum(other int) int {
	if int(i) > other {
		return int(i)
	}
	return other
}

func ReturnLargestNum(x, y, z int) (largest int) {
	largest = MyInt(x).ReturnLargeNum(y)
	if largest < MyInt(y).ReturnLargeNum(z) {
		largest = MyInt(y).ReturnLargeNum(z)
	}
	if largest < MyInt(z).ReturnLargeNum(x) {
		largest = MyInt(z).ReturnLargeNum(x)
	}

	return
}
