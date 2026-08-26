package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _,num := range nums{
		total += num
	}

	return total
}

func main() {
    nums := []int{1,3,2,4}
	fmt.Println(sum(nums...))
}
