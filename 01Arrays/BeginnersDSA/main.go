package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// 1. Two Sum - 20 Jan 2022
	var unsortedEle = []int{9, 0, 5, 8, 11, 1}
	fmt.Println("Two sum positive result of Unsorted list: ", twoSumUnsortedArr(unsortedEle, 10))
	fmt.Println("Two sum negetive result of Unsorted list: ", twoSumUnsortedArr(unsortedEle, 15))

	var sortedEle = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Two sum positive result of sorted list: ", twoSumSortedArr(sortedEle, 10))
	fmt.Println("Two sum negetive result of sorted list: ", twoSumSortedArr(sortedEle, 13))

	// 2. Best Time to Buy and Sell Stock
	var pricesTestCase1 = []int{7, 1, 5, 3, 6, 4}
	fmt.Println("TC1 - Max Profit: ", maxProfitBuyAndSellStocks(pricesTestCase1))
	var pricesTestCase2 = []int{7, 6, 4, 3, 1}
	fmt.Println("TC2 - Max Profit: ", maxProfitBuyAndSellStocks(pricesTestCase2))

	// 3. Merge 2 Sorted arrays
	var mergeArr1 = []int{1, 3, 5, 7, 8}
	var mergeArr2 = []int{2, 4, 9, 11}
	fmt.Println("Result of merging 2 sorted arrays is: ", mergeTwoSortedArrays(mergeArr1, mergeArr2))

	// 4. Move all zeros to begining
	var moveZeroArr = []int{1, 1, 0, 2, 0, 4, 5}
	fmt.Println("Result of moving zeros to begining is: ", moveZerosToBegining(moveZeroArr))

	// 5. Find duplicate in an unsorted array
	var findDuplicatesInUnsorted = []int{6, 3, 1, 5, 4, 3, 2}
	fmt.Println("Result of duplicate ele in unsorted array is: ", duplicateInUnsortedArray(findDuplicatesInUnsorted))

	// 6. Find majority element
	var findMajorityElement1 = []int{2, 3, 9, 2, 2}
	var findMajorityElement2 = []int{8, 5, 1, 9}
	fmt.Println("Majority lement for case1: ", findMajorityElement(findMajorityElement1))
	fmt.Println("Majority lement for case2: ", findMajorityElement(findMajorityElement2))

	// 7. Occurs Once
	var findOccursOnce = []int{2, 3, 2, 3, 6}
	fmt.Println("Ele occuring once is: ", occursOnce(findOccursOnce))

	// 8. Jewels and Stones
	var jewels string = "aA"
	var stones string = "aAAbbAbcAa" // Ans: 6
	fmt.Println("Count of jewels in stone: ", jewelsAndStones(jewels, stones))

	// 9. Sorting the sentence
	var s string = "is2 sentence4 This1 a3"
	fmt.Println("Result of sorting a sentence is: ", sortSentece(s))

	// 10. Uppercase to lowercase
	var uppercaseToLowerCase string = "Hello"
	fmt.Println("Uppercase to lowercase: ", uppercaseToLowercase(uppercaseToLowerCase))
}

func maxProfitBuyAndSellStocks(prices []int) int {

	minPrice := math.MaxInt64
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

func twoSumSortedArr(arr []int, target int) []int {
	first := 0
	last := len(arr) - 1
	result := make([]int, 0)

	for first < last {
		potentialSum := arr[first] + arr[last]
		if potentialSum == target {
			result = append(result, arr[first])
			result = append(result, arr[last])
			return result
		} else if potentialSum > target {
			last--
		} else if potentialSum < target {
			first++
		}
	}
	result = append(result, -1)
	return result
}

func twoSumUnsortedArr(arr []int, target int) []int {
	maps := make(map[int]int)
	result := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		potentialEle := target - arr[i]
		_, ok := maps[potentialEle]

		if ok {
			result = append(result, potentialEle)
			result = append(result, arr[i])
			return result
		}
		maps[arr[i]] = i
	}
	result = append(result, -1)
	return result
}

func mergeTwoSortedArrays(arr1 []int, arr2 []int) []int {
	i := 0
	j := 0
	arr3 := make([]int, 0)
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		arr3 = append(arr3, arr1[i])
		i++
	}

	for j < len(arr2) {
		arr3 = append(arr3, arr2[j])
		j++
	}
	return arr3
}

func moveZerosToBegining(arr []int) []int {
	count := len(arr) - 1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count--
		}
	}
	for count >= 0 {
		arr[count] = 0
		count--
	}
	return arr
}

func duplicateInUnsortedArray(arr []int) int {
	ele := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		_, ok := ele[arr[i]]
		if ok {
			return arr[i]
		} else {
			ele[arr[i]] = i
		}
	}
	return -1
}

func findMajorityElement(arr []int) int {
	majorityEle := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		_, ok := majorityEle[arr[i]]
		if ok {
			majorityEle[arr[i]] = majorityEle[arr[i]] + 1
		} else {
			majorityEle[arr[i]] = 1
		}
	}
	max := 1
	ele := -1
	for key, val := range majorityEle {
		if val > max {
			ele = key
			max = val
		}
	}
	return ele
}

func occursOnce(arr []int) int {
	ele := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		val, ok := ele[arr[i]]
		if ok {
			ele[arr[i]] = val + 1
		} else {
			ele[arr[i]] = 1
		}
	}
	for key, val := range ele {
		if val == 1 {
			return key
		}
	}
	return 0
}

func jewelsAndStones(jewels string, stones string) int {
	jewelMap := make(map[string]int)
	for i := 0; i < len(jewels); i++ {
		_, ok := jewelMap[string(jewels[i])]
		if !ok {
			jewelMap[string(jewels[i])] = 1
		}
	}
	count := 0
	for i := 0; i < len(stones); i++ {
		_, ok := jewelMap[string(stones[i])]
		if ok {
			count = count + 1
		}
	}
	return count
}

func sortSentece(sentence string) string {
	result := make([]string, 8)
	senArr := strings.Split(sentence, " ")
	for _, val := range senArr {
		lastChar, _ := strconv.ParseInt(val[len(val)-1:], 0, 32)
		result[lastChar-1] = val[:len(val)-1]
	}
	ret := strings.Join(result, " ")

	return ret
}

func uppercaseToLowercase(str string) string {
	uppercaseStart := 65
	lowerCaseStart := 97
	for i := 0; i < len(str); i++ {
		if int(str[i]) >= uppercaseStart && int(str[i]) <= 90 {
			diffNumber := int(str[i]) - uppercaseStart
			lowercaseASCII := lowerCaseStart + diffNumber
			lowercaseChar := string(lowercaseASCII)
			str = strings.Replace(str, string(str[i]), lowercaseChar, -1)
		}
	}
	return str
}
