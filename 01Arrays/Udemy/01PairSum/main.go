package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1. Pair Sum
	arr := []int{10, 5, 2, 3, -6, 9, 11}
	s := 4

	p := pairSum(arr, s)

	if len(p) == 0 {
		fmt.Println("No such pair")
	} else {
		fmt.Println(p)
	}

}

func pairSum(arr []int, targetSum int) []int {
	s := make(map[int]int)
	result := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		potentialEle := targetSum - arr[i]
		_, ok := s[potentialEle]

		if !ok {
			result = append(result, potentialEle)
			result = append(result, arr[i])
			break
		}
		s[arr[i]] = 1
	}
	sort.Ints(result)
	return result
}
