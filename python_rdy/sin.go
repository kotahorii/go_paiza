package pythonrdy

import "math"

func CalcSin(degree int) float64 {
	return math.Round(math.Sin(float64(degree)*math.Pi/180)*10) / 10
}
