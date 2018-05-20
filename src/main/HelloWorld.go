package main

import "fmt"

func main()  {
	fmt.Println("Hello World!")

	for i, v := range "Go语言" {
		fmt.Printf("%d: %c\n", i, v)
	}

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8] // []int{5,6}
	length := (2)
	capacity := (4)
	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	slice5 = slice5[:cap(slice5)] // []int{5,6,7,8}
	slice5 = append(slice5, 11, 12, 13) // []int{5,6,7,8,11,12,13}
	length = (7)
	fmt.Printf("%v\n", length == len(slice5))
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6) // []int{0,0,0,8,11,12,13}
	e2 := (0)
	e3 := (8)
	e4 := (11)
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])

	f := func(i int) int {
		fmt.Printf("%d ",i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}

}





