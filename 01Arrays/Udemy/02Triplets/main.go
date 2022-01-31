package main

import "sort"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}
	s := 18

	triplets(arr, s)

}

func triplets(arr []int, s int) {

	sort.Ints(arr)

	// Pick every a[i], and solve pair sum for remaining part
	for i := 0; i < len(arr)-3; i++ {
		// j := i +1
		// k :=
	}

	// result := make([][]int,)
}
