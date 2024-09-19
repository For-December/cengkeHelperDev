package generator

import "cengkeHelperDev/src/utils/logger"

func Bin2WeekLesson(binNum uint32) ([]int, []int) {
	weekNums := make([]int, 0)
	lessonNums := make([]int, 0)

	for i := 1; i <= 19; i++ {
		if (1<<(32-i))&binNum == 0 {
			continue
		}
		weekNums = append(weekNums, i)
	}

	for i := 1; i <= 13; i++ {
		if (1<<(i-1))&binNum == 0 {
			continue
		}
		lessonNums = append(lessonNums, i)
	}

	return weekNums, lessonNums

}

func WeekLesson2Bin(weekNums, lessonNums []int) uint32 {
	var res uint32 = 0
	for _, num := range weekNums {
		if num > 19 {
			logger.Error("weekNum不能超过19", num)
		}
		if res&((1<<31)>>(num-1)) != 0 {
			logger.Error("weekNum重复覆盖！")
		}
		res = res | ((1 << 31) >> (num - 1))
	}

	for _, num := range lessonNums {
		if num > 13 {
			logger.Error("lessonNums不能超过13", num)
		}
		if res&(1<<(num-1)) != 0 {
			logger.Error("lessonNum重复覆盖！")
		}
		res = res | (1 << (num - 1))
	}

	return res
}
