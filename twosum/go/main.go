package main

import (
	"fmt"
)

func main() {
	ex1 := []int{2, 7, 11, 15}
	ex2 := []int{3, 2, 4}
	ex3 := []int{3, 3}

	fmt.Printf("result 1: %v \n", twoSum(ex1, 9))
	fmt.Printf("result 2: %v \n", twoSum(ex2, 6))
	fmt.Printf("result 3: %v \n", twoSum(ex3, 6))
}

func twoSum(nums []int, target int) []int {
	out := []int{}
	exist := map[int]int{}
	for i, num := range nums {
		j, found := exist[target-num]
		if found {
			out = append(out, j, i)
			return out
		}

		exist[num] = i
	}

	return out
}
