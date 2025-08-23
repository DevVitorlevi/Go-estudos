package main

import "fmt"

func main() {
	// Arrays
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	//Slices
	slice := []int{1, 3, 4, 4, 5, 6}
	slice = append(slice, 2)
	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)
}
