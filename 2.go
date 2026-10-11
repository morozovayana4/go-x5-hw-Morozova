package main

import "fmt"

func main() {
	var numbers []int
	var sum int
	var max int
	var min int
	var sred float64
	min = 1e12
	max = -1
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		sum += num
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		numbers = append(numbers, num)
	}
	if len(numbers) == 0 {
		fmt.Println("input is empty")
	} else {
		sred = float64(sum) / float64(len(numbers))
		fmt.Printf("sum = %d, max = %d, min = %d, sred = %.2f\n", sum, max, min, sred)
	}
}
