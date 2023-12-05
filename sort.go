package main

import (
	"fmt"
	"time"
)

func BubbleSort(arr *[]int) []int {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	return *arr
}

func TestBubbleSort(arr []int) []int {
	now := time.Now()
	sortedArr := BubbleSort(&arr)
	fmt.Printf("Optimized Time taken to Bubble sort an array of length %v : %v\n", len(arr), time.Since(now))
	return sortedArr
}
