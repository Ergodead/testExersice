package main

import "fmt"

func uniqueElementsOfSlice(arr []int) []int {
	mapOfElements := make(map[int]bool)
	resultSlice := []int{}
	for _, x := range arr {
		if !mapOfElements[x] {
			mapOfElements[x] = true
			resultSlice = append(resultSlice, x)
		}

	}
	return resultSlice
}

func cross(arrA, arrB []int) []int {
	mapOfSliceA := make(map[int]bool)
	for _, x := range arrA {
		mapOfSliceA[x] = true
	}
	result := []int{}
	mapOfSliceB := make(map[int]bool)
	for _, x := range arrB {
		if mapOfSliceA[x] && !mapOfSliceB[x] {
			result = append(result, x)
			mapOfSliceB[x] = true
		}
	}
	return result
}

func summUniqueElements(arrA, arrB []int) []int {
	summ := append(arrA, arrB...)
	return uniqueElementsOfSlice(summ)
}

func main() {
	a := []int{1, 4, 2, 5, 2, 7}
	b := []int{4, 7, 9, 1, 6, 4, 3}

	uniqueA := uniqueElementsOfSlice(a)
	uniqueB := uniqueElementsOfSlice(b)
	crossing := cross(a, b)
	summ := summUniqueElements(a, b)

	fmt.Println(uniqueA, uniqueB)
	fmt.Println(crossing)
	fmt.Println(summ)
}
