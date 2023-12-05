package main

import (
	"math"
)


func linearSearch(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if target > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func SearchInt(arr []int, val int, ordered bool) int {
	if ordered {
		if len(arr) > 10000 {
			return TestBinarySearch(arr, val)
		} else {
			return TestLinearSearch(arr, val)
		}
	}
	return TestLinearSearch(arr, val)
}

func BinaryCrystalBall(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount
	for j := 0; j < jumpAmount && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
		i++
	}

	return -1

}

