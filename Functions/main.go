package main

import "fmt"

func main() {
	//numbers := []int{1, 10, 15}
	sum := sumUp(1, 10, 15)

	fmt.Println(sum)
}

func sumUp(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
