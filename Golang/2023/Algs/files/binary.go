package algs

import (
	"fmt"
)

func BinarySearch() {
	fmt.Print("Binary search.\n")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 19, 122}
	var target int

	fmt.Scanf("%d", &target)
	low := 0
	max := len(arr) - 1

	for low <= max {
		mid := (low + max) / 2
		guess := arr[mid]
		fmt.Printf("target %d, mid %d, guess %d, low %d, max %d\n", target, mid, guess, low, max)
		if guess == target {
			fmt.Print(guess)
			return
		}
		if guess > target {
			max = mid - 1
		}
		if guess < target {
			low = mid + 1
		}
	}
	fmt.Println(arr[low])
}
