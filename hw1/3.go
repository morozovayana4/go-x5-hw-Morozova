package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scan(&a)
	if err == nil && a >= 0 {
		if a > 50000 {
			fmt.Println("large")
		} else if a < 10000 {
			fmt.Println("small")
		} else {
			fmt.Println("regular")
		}
	} else {
		fmt.Println("input error")
	}
}
