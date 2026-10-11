package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scan(&a)
	if err == nil && a > 0 && a <= 30 {
		if a%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if a%3 == 0 {
			fmt.Println("Fizz")
		} else if a%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(a)
		}
	} else {
		fmt.Println("input error")
	}
}
