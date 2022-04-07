package pythonrdy

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ex 5C => 41F
func ConvTemperature(T string) (convertedT string) {
	// divide temperature and unit
	rexTemperature := regexp.MustCompile("[0-9]+")
	rexUnit := regexp.MustCompile("[a-zA-Z]+")

	temperature, err := strconv.Atoi(rexTemperature.FindString(T))
	if err != nil {
		log.Fatal(err)
	}
	unit := rexUnit.FindString(T)

	// convert temperature
	unit = strings.ToUpper(unit)
	if unit == "C" {
		convertedT = fmt.Sprintf("%.1fF", float64(temperature)*9/5+32)
	} else if unit == "F" {
		convertedT = fmt.Sprintf("%.1fC", float64(temperature-32)*5/9)
	} else {
		log.Fatal("invalid unit")
	}
	return
}
