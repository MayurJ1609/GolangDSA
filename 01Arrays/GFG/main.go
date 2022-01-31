package main

import "fmt"

func main() {

	// 1. Find duplicates in an array
	findDuplicatesArr := []int{2, 3, 1, 2, 3}
	fmt.Println(findDuplicates(findDuplicatesArr))

	// 2. Find transition point
	transitionArr := []int{0, 0, 0, 1, 1}
	fmt.Println("Trasition index for transitionArr is: ", trasitionPoint(transitionArr))

	// 3. Peak element

	// 4. Remove duplicate elements from sorted Array
	a := []int{1, 1, 2, 2}
	retVal := removeDuplicateMethodOne(a, len(a))
	fmt.Println("MethodOne Remove duplicates: ", retVal)
	removeDuplicateMethodTwo(a, len(a))

}

// 1. Find duplicates in an array
func findDuplicates(arr []int) []int {

	duplicateMap := make(map[int]int)
	retArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		_, ok := duplicateMap[arr[i]]
		if ok {
			retArr = append(retArr, arr[i])
		} else {
			duplicateMap[arr[i]] = i
		}
	}
	if len(retArr) == 0 {
		retArr = append(retArr, -1)
	}
	return retArr
}

// Remove duplicates Method 1
func removeDuplicateMethodOne(a []int, n int) []int {

	duplicateMap := make(map[int]int)
	retVal := make([]int, 0)

	for i := 0; i < len(a); i++ {
		_, ok := duplicateMap[a[i]]
		if !ok {
			retVal = append(retVal, a[i])
			duplicateMap[a[i]] = a[i]
		}
	}
	return retVal
}

// Remove duplicates Method 2
func removeDuplicateMethodTwo(a []int, n int) {
	j := 0
	for i := 0; i < n-1; i++ {
		if a[i] != a[i+1] {
			a[j] = a[i]
			j++
		}
	}
	a[j] = a[n-1]
	j++
	fmt.Println("MethodOne Remove duplicates: ", a[:j])
}

func trasitionPoint(a []int) int {

	start := 0
	end := len(a) - 1

	for start <= end {
		mid := (start + end) / 2
		if a[mid] < a[end] {
			start = mid + 1
		} else if a[mid] > a[start] {
			end = mid + 1
		} else if a[mid] == 1 {
			return mid
		}
	}

	return -1
}
