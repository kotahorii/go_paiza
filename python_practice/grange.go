package pythonpractice

func Grange(lower, upper, step int) []int {
	if upper == lower {
		return []int{}
	}

	if (step > 0 && lower > upper) || (step < 0 && lower < upper) {
		return []int{}
	}

	return addLowerToUpper(lower, upper, step)
}

func addLowerToUpper(lower, upper, step int) (result []int) {
	if upper > lower {
		for i := lower; i < upper; i += step {
			result = append(result, i)
		}
	} else if upper < lower {
		for i := lower; i > upper; i += step {
			result = append(result, i)
		}
	}
	return
}
