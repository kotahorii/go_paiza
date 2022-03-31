package question

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func DisplaySPAM10Times() []string {
	var result []string
	for i := 0; i < 10; i++ {
		result = append(result, "SPAM")
	}
	return result
}

func DisplayThreeTimes() {
	for i := 1; i < 10; i++ {
		fmt.Printf("3 * %d = %d\n", i, 3*i)
	}
}

func DisplayTwoToThePower() {
	for i := 1; i < 9; i++ {
		fmt.Printf("2 to the power of %d = %d\n", i, int(math.Pow(2, float64(i))))
	}
}

func ReturnFactorial(n int) (result int) {
	result = 1
	for i := 1; i < n+1; i++ {
		result *= i
	}
	return
}

func CalcMeanValue() (mean int) {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 10; i++ {
		if sc.Scan() {
			num, err := strconv.Atoi(sc.Text())
			if err != nil {
				log.Fatal(err)
			}
			mean += num
		}
	}

	mean /= 10
	return
}

func CountWinAndLose() (win, lose int) {
	count := 0
	sc := bufio.NewScanner(os.Stdin)

	for count < 10 {
		fmt.Printf("%d times: ", count+1)
		if sc.Scan() {
			if num, err := strconv.Atoi(sc.Text()); err != nil {
				log.Println("invalid: ", err)
			} else {
				switch num {
				case 1:
					win += 1
					count += 1
				case 0:
					lose += 1
					count += 1
				default:
					fmt.Println("please input 1 or 0")
				}
			}
		}
	}

	return
}

func DisplayBaseBallScore() {
	count := 0
	sc := bufio.NewScanner(os.Stdin)
	kyojinScore := 0
	hanshinScore := 0

	for count < 9 {
		// kyojin score
		fmt.Printf("%d front: ", count+1)
		if sc.Scan() {
			if num, err := strconv.Atoi(sc.Text()); err != nil {
				log.Println("invalid: ", err)
				continue
			} else {
				kyojinScore += num
			}
		}

		// hanshin score
		fmt.Printf("%d back: ", count+1)
		if sc.Scan() {
			if num, err := strconv.Atoi(sc.Text()); err != nil {
				log.Println("invalid: ", err)
				continue
			} else {
				hanshinScore += num
			}
		}

		count += 1
	}
	fmt.Printf("kyojin: %d, hanshin: %d\n", kyojinScore, hanshinScore)
	if kyojinScore < hanshinScore {
		fmt.Println("hanshin won")
	} else if kyojinScore > hanshinScore {
		fmt.Println("Kyojin won")
	} else {
		fmt.Println("draw")
	}
}
