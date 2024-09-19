package main

import "fmt"

func Test() {
	fmt.Println(
		bin2WeekLesson(
			weekLesson2Bin([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 19, 18, 17}, []int{3, 7, 2, 4, 13, 11, 12}),
		),
	)
}
