package task_1

import "fmt"

func Run() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//a = delete(a, 1)
	fmt.Println(a)
	fmt.Println(indexof(a, func(i int) bool {
		return i == 5
	}))
	fmt.Println(indexof(a, func(i int) bool {
		return i == 20
	}))
	fmt.Println(some(a, func(i int) bool {
		return i == 5
	}))
	fmt.Println(some(a, func(i int) bool {
		return i == 20
	}))
}
