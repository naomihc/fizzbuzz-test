package main

import "fmt"

func isDivBy3(num int) bool {
	return num%3 == 0
}

func isDivBy5(num int) bool {
	return num%5 == 0
}

func main() {
	for i := 0; i < 100; i++ {
		if isDivBy3(i) {
			if isDivBy5(i) {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if isDivBy5(i) {
			fmt.Println("Buzz")
		}
	}
}
