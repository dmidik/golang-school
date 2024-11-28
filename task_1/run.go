package task_1

import "fmt"

func Run() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	a = delete(a, 1)
	fmt.Println(a)
}
