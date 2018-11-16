
package main

import (
	"fmt"
	"math"
)

func main()  {
	grades := []float64{ 1, 2, 3, 4 }
	avg := Average(grades)
	fmt.Println(avg)
	fmt.Println(Max(grades))
	fmt.Println(Min(grades))

}

func Average(numbers []float64) (average float64)  {
	if len(numbers) == 0 {
		return 0
	}
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	average = total / float64(len(numbers))
	return

}

func Min(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	min := math.MaxFloat64
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min

}

func Max(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	max := -math.MaxFloat64
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max

}