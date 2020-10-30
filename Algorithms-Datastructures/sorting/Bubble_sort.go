package main

import (
	"fmt"
	_ "os"
)

func main() {
	a := []int{31, 8, 6, 54, 95, 84, 71, 67}
	fmt.Printf("%v\n", a)

	len_arr := len(a) - 1
	fmt.Printf("The length of array : %v\n", len_arr)

	for i := 0; i < len_arr; i++ {
		for j := 0; j < len_arr-i; j++ {
			fmt.Printf("first index:%v The value : %v\n", a[j], a[j+1])
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp

			}
		}
		fmt.Println("----Iteration----", i+1)

		fmt.Printf("\nYour sort Array : %v\n", a)

	}

	fmt.Printf("Final sorted Array : %v", a)
}
