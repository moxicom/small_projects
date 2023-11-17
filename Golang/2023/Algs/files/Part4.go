package algs

import (
	"fmt"
	"time"
)

func RecursiveSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + RecursiveSum(arr[1:])
}

func ArrayLen(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}
	return 1 + ArrayLen(arr[1:])
}

func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func QuickSort(arr []int) {
	n := len(arr)

	if n < 2 {
		return
	}
	if n == 2 {
		if arr[0] > arr[1] {
			swap(arr, 0, 1)
		}
	}

	pivot := n - 1
	rmi := 0
	for i := 0; i < pivot; i++ {
		if arr[i] <= arr[pivot] {
			swap(arr, i, rmi)
			rmi++
		}
	}
	QuickSort(arr[:rmi])
	QuickSort(arr[rmi:])
}

func summarizing() {
	arr := []int{10, 2, 4, 3, 4}
	fmt.Println("Start: ", time.Now().Local())
	QuickSort(arr)
	fmt.Println(arr)
	fmt.Println("End: ", time.Now().Local())
}
