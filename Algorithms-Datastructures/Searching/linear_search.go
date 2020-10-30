package main

import "fmt"

func main() {
	arr := []int{87, 58, 46, 78, 35, 86, 19, 21, 20}
	a := linearsearch(arr, 35)
	fmt.Printf("The Source Array : %v\n", arr)
	fmt.Printf("The element %v is found at %v location", 35, a)
}

func linearsearch(arr []int, search int) int {
	for i, item := range arr {
		if item == search {
			return i + 1
		}
	}
	return 0
}
