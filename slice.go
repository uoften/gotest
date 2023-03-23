package main

import "fmt"

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func main()  {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a[3:4])
	fmt.Println(a[3:4:4])
	//
	//var array [10]int
	//var slice = array[5:6]
	//fmt.Println(slice[0])
	//fmt.Println("lenth of slice: ", len(slice))
	//fmt.Println("capacity of slice: ", cap(slice))
	//fmt.Println(&slice[0] == &array[5])

	var slice2 []int
	for i:=1;i<=40;i++ {
		slice2 = append(slice2, i)
		if len(slice2) < cap(slice2) {
			//fmt.Println(len(slice2))
			fmt.Println(cap(slice2))
			//fmt.Println(&slice2[0])
		}
	}

	//slice2 = append(slice2, 1)
	//slice2 = append(slice2, 2)
	//slice2 = append(slice2, 3)

	slice2 = append(slice2, 1,2,3,4,5,6,7)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	newSlice := AddElement(slice2, 13)
	fmt.Println(&slice2[0])
	fmt.Println(&newSlice[0])
	fmt.Println(&slice2[0] == &newSlice[0])
}
