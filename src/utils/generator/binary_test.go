package generator

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(
		Bin2WeekLesson(
			WeekLesson2Bin([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 19, 18, 17},
				[]int{3, 7, 2, 4, 13, 11, 12}),
		),
	)

	result1, result2 := Bin2WeekLesson(
		WeekLesson2Bin([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 19, 18, 17},
			[]int{3, 7, 2, 4, 13, 11, 12}),
	)

	if len(result1) != 12 || len(result2) != 7 {
		t.Error("error")
	}

	for i := 1; i <= 19; i++ {
		if res1, res2 := Bin2WeekLesson(
			WeekLesson2Bin([]int{i},
				[]int{13})); res1[0] != i || res2[0] != 13 {
			t.Error("周出错")
		}
	}

	for i := 1; i <= 13; i++ {
		if res1, res2 := Bin2WeekLesson(
			WeekLesson2Bin([]int{19},
				[]int{i})); res1[0] != 19 || res2[0] != i {
			t.Error("节出错")
		}
	}
	//t.Errorf("Add(3, 4) = %d; want 7", result)

}
