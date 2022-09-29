package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 || args[1] != "data.txt" {
		fmt.Println("ERROR: invalid number of arguments")
		return
	}
	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println("ERROR: read file")
		return
	}
	byteData := strings.Fields(string(data))
	var numbers []float64
	for _, number := range byteData {
		convert, err := strconv.ParseFloat(number, 64)
		if err != nil {
			fmt.Println("ERROR: converting number")
			return
		}
		numbers = append(numbers, convert)
	}

	n := len(numbers)

	var sumX, sumY, sumX2, sumY2, sumXY float64

	for x, y := range numbers {
		sumX += float64(x)
		sumY += float64(y)
		sumX2 += float64(x * x)
		sumY2 += float64(y * y)
		sumXY += float64(x * int(y))
	}
	a := (sumY*sumX2 - sumX*sumXY) / (float64(n)*(sumX2) - math.Pow((sumX), 2))
	b := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - math.Pow(sumX, 2))
	r := (float64(n)*sumXY - sumX*sumY) / (math.Pow((float64(n)*sumX2-math.Pow(sumX, 2))*(float64(n)*sumY2-math.Pow(sumY, 2)), 0.5))

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", b, a)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
