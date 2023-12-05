package main

import "math/rand"

func getRandomArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size * 10)
	}
	return arr
}

func printArr(arr []int) {
	print("[")
	for i, v := range arr {
		if i == len(arr)-1 {
			print(v)
		} else {
			print(v, ",")
		}
	}
	println("]")
	println("length of array:", len(arr))
}
