package main

import "fmt"

func main() {
	// // Shorthand declaration of array
	// arr1 := [4]int{1, 2, 3}
	// arr2 := [...]int{10, 20, 30}
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// // LongHand declaration of array
	// var arr3 [3]int = [3]int{11, 12, 13}
	// fmt.Println(arr3)

	// var arr4 [3]int
	// arr4[0] = 111
	// fmt.Println(arr4)

	// arr_dimetion := [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println(arr_dimetion)

	// for loop on array
	// for x := 0; x < 3; x++ {
	// 	fmt.Println(arr3[x])
	// }

	// for x := 0; x < 3; x++ {
	// 	for y := 0; y < 3; y++ {
	// 		fmt.Println(arr_dimetion[x][y])
	// 	}
	// }

	// copy array to another variable
	// in golang : arrays are not coppied by reference but values

	a := [...]int{1, 2, 3}
	b := a
	b[0] = 100
	//  a[0] and  b[0] : both are different
	fmt.Println(a)
	fmt.Println(b)

}