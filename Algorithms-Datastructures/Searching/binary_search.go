package main

import (
	"fmt"
	_ "os"
)

func main() {
	a := []int{6, 8, 31, 54, 67, 71, 84, 95}
	fmt.Printf("%v", a)
	fmt.Println("\nPlease enter search value?")
	var search_elem int
	fmt.Scanf("%d", &search_elem)
	fmt.Println("Read Number", search_elem)

	var mid int
	start := 0

	len_arr := len(a) - 1
	end := len_arr

	for start <= end {
		mid = (start + end) / 2
		if search_elem == a[mid] {
			fmt.Printf("%d found at location %d.\n", search_elem, mid+1)
			break
		} else if search_elem > a[mid] {
			start = mid + 1
		} else if search_elem < a[mid] {
			end = mid - 1
		}

	}

	if start > end {
		fmt.Printf("%d is not present in the list.\n", search_elem)
	}
}
