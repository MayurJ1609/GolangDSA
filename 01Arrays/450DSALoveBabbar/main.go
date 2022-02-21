package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Welcome to love Babbar sheet of 450 DSA problems")

	// 21-02-22 Array Subset of another Array
	var a1 = []int{11, 1, 13, 21, 3, 7}
	var a2 = []int{11, 3, 7, 1}
	fmt.Println("Subset: ", arraySubset(a1, a2))

	var b1 = []int{10, 5, 2, 23, 19}
	var b2 = []int{19, 5, 3}
	fmt.Println("Subset: ", arraySubset(b1, b2))

	// 21-02-22 Chocolate distribution problem
	var chocolateDistributionArr = []int{3, 4, 1, 9, 56, 7, 9, 12}
	fmt.Println("Chocolate distribution: ", chocolateDistribution(chocolateDistributionArr, 5))

	// 21-02-22 Three way partition
	var threewayPartitionArr = []int{1, 2, 3, 3, 4}
	fmt.Println("Threeway partition: ", threewayPartition(threewayPartitionArr, 1, 2))

}

func arraySubset(a1 []int, a2 []int) bool {

	hm := make(map[int]int) // space O(n)
	for i := 0; i < len(a1); i++ {
		_, ok := hm[a1[i]]
		if !ok {
			hm[a1[i]] = i
		}
	}

	for i := 0; i < len(a2); i++ {
		_, ok := hm[a2[i]]
		if !ok {
			return false
		}
	}
	return true
}

func chocolateDistribution(a []int, m int) int {

	sort.Ints(a)
	min := math.MaxInt

	for i := 0; i+m-1 < len(a); i++ {
		d := a[i+m-1] - a[i]

		if d < min {
			min = d
		}
	}
	return min

}

func threewayPartition(array []int, a int, b int) {
	l := 0
	r := len(array) - 1

	for i := 0; i < r; i++ {
		if array[i] < a {
			temp := array[i]
			array[i] = array[l]
			array[l] = temp
			l++
		} else if array[i] > b {
			temp := array[i]
			array[i] = array[r]
			array[r] = temp
			r--
			i--
		}
	}
}
