package main

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/generator"
	"cengkeHelperDev/src/utils/logger"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var count1 = 0
var count2 = 0
var count3 = 0

// 节号相同，如3-3节
var count4 = 0

// 匹配单条记录
var re = regexp.MustCompile(`\S+?\(\d+?,\S+?\)`)

func getBuilding(str string) (string, string) {
	splits := strings.Split(str, "-")
	if len(splits) > 2 {
		logger.Error("授课地点切分后大于2,有误: ", str)
	}

	curBuilding := splits[0]
	if len(splits) == 2 {
		curBuilding += "-教学楼"
		return curBuilding, splits[1]
	}

	// 不能切分,可能是字母或者其他
	curBuilding = string([]rune(str)[0])
	if curBuilding == "A" {
		curBuilding += "-教学楼"
		return curBuilding, string([]rune(str)[1:])
	}
	return str, str

}

func getRoom(str string) (string, string) {
	splits := strings.Split(str, "(")
	if len(splits) > 2 {
		logger.Error("授课地点切分后大于2,有误: ", str)
	}
	return splits[0], splits[1]
}

func getAreaNums(str string) (int, string) {
	areaNum := 0
	// 提取区号
	idx := strings.Index(str, "区")
	if idx != -1 {
		tempStr := str[:idx]
		areaNum, _ = strconv.Atoi(tempStr)
	}

	if areaNum == 0 {
		logger.Error("areaNum error")
	}

	return areaNum, str[idx+len("区"):]

}
func getTimeNums(str string, splitWord string) (res []int, newStr string) {

	res = make([]int, 0)
	for idx := strings.Index(str, splitWord); idx != -1; {
		tempStr := strings.Trim(str[:idx], ",")

		// 去掉已经提取的部分，实现继续提取
		str = str[idx+len(splitWord):]
		idx = strings.Index(str, splitWord)

		newStr = str

		// 没找到分隔符则返回整个切片
		split := strings.Split(tempStr, "-")
		if len(split) < 1 || len(split) > 2 {
			logger.Error("split error")
			return
		}
		// 1或2个数字
		startNum, err := strconv.Atoi(split[0])
		if err != nil {
			logger.Error(err)
			return
		}
		// 加入第一个数字
		res = append(res, startNum)

		if len(split) == 1 {
			continue
		}

		// 加入第一到第二个数字之间的所有数字
		endNum, err := strconv.Atoi(split[1])
		if err != nil || startNum == endNum {
			logger.Warning("第二个数字和第一个一样：", newStr, err)
			count4++
			continue
		}

		for i := startNum + 1; i <= endNum; i++ {
			res = append(res, i)
		}

	}

	return
}

var timeInfoIdCount = 0

func getCellTimeInfo(infoStr string, curWeekday int, courseInfoId uint32) dbmodels.TimeInfo {
	infoStr = strings.Trim(infoStr, ",")

	if strings.Contains(infoStr, "操场") ||
		strings.Contains(infoStr, "体育馆") {
		// 操场和体育馆单独处理
		return dbmodels.TimeInfo{}
	}

	// 提取周号
	weekNums, newStr := getTimeNums(infoStr, "周")
	// 提取节号(基于周号继续提取)
	lessonNums, newStr2 := getTimeNums(newStr, "节")

	// 提取区号
	areaNum, newStr3 := getAreaNums(newStr2)

	fmt.Println(weekNums, lessonNums, areaNum, infoStr)

	// 提取教学楼
	building, newStr4 := getBuilding(newStr3)

	// 提取教室
	room, _ := getRoom(newStr4)
	fmt.Println(building, room)

	timeInfoIdCount++
	return dbmodels.TimeInfo{
		ID:           uint32(timeInfoIdCount),
		CourseInfoId: courseInfoId,
		WeekAndTime:  generator.WeekLesson2Bin(weekNums, lessonNums),
		DayOfWeek:    uint8(curWeekday),
		Area:         uint8(areaNum),
		Building:     building,
		Classroom:    room,
	}

}
func getAddrTimeInfo(timeInfos *[]dbmodels.TimeInfo,
	timeAndAddr string, curWeekday int, courseInfoId uint32) {
	if timeAndAddr == "" {
		return
	}
	if !strings.Contains(timeAndAddr, "区") ||
		strings.Contains(timeAndAddr, "(单)") ||
		strings.Contains(timeAndAddr, "(双)") ||
		strings.Contains(timeAndAddr, "节(,)") {
		return
	}

	count2++

	matches := re.FindAllStringSubmatch(timeAndAddr, -1)

	//// 有多个‘周’字的结果
	//if len(matches) != len(strings.Split(timeAndAddr, ","))-1-len(matches)+1 {
	//
	//	println("> ", len(matches), len(strings.Split(timeAndAddr, ","))-1,
	//		matches[0][0], " <-> ", timeAndAddr)
	//} else {
	//
	//	// 只有单个‘周’字
	//}
	if matches == nil || len(matches) == 0 {
		println(timeAndAddr)
		os.Exit(2)
	}

	for _, match := range matches {
		count3++
		*timeInfos = append(*timeInfos,
			getCellTimeInfo(match[0], curWeekday, courseInfoId))
	}

	//println(timeAndAddr)
}

func UpdateDB() {

	//return
	info, err := generator.ReadTeachInfo()
	if err != nil {
		logger.Error(err)
		return
	}

	resCourse := make([]dbmodels.CourseInfo, 0)
	resTime := make([]dbmodels.TimeInfo, 0)

	//count := 0
	for i, teachInfo := range info {
		count1++
		resCourse = append(resCourse, dbmodels.CourseInfo{
			ID:               uint32(i + 1),
			Years:            teachInfo.Years,
			Semester:         teachInfo.Semester,
			CourseNum:        teachInfo.CourseNum,
			CourseName:       teachInfo.CourseName,
			Faculty:          teachInfo.Faculty,
			Credit:           teachInfo.Credit,
			CourseComplexion: teachInfo.CourseComplexion,
			CourseType:       teachInfo.CourseType,
			Grade:            teachInfo.Grade,
			Major:            teachInfo.Major,
			Teacher:          teachInfo.Teacher,
			TeacherTitle:     teachInfo.TeacherTitle,
		})

		getAddrTimeInfo(&resTime, teachInfo.AddrSunday, 0, uint32(i+1))
		getAddrTimeInfo(&resTime, teachInfo.AddrMonday, 1, uint32(i+1))
		getAddrTimeInfo(&resTime, teachInfo.AddrTuesday, 2, uint32(i+1))
		getAddrTimeInfo(&resTime, teachInfo.AddrWednesday, 3, uint32(i+1))
		getAddrTimeInfo(&resTime, teachInfo.AddrThursday, 4, uint32(i+1))
		getAddrTimeInfo(&resTime, teachInfo.AddrFriday, 5, uint32(i+1))
		getAddrTimeInfo(&resTime, teachInfo.AddrSaturday, 6, uint32(i+1))

	}
	//logger.Info("start")
	//println(database.Client.Error)
	//if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
	//	logger.Error("run server error: ", err)
	//	return
	//}

	println(count1, count2, count3, count4)
	println(len(resCourse), len(resTime))

	batchInsert[dbmodels.CourseInfo](resCourse)
	batchInsert[dbmodels.TimeInfo](resTime)

	logger.Info("更新数据完毕")

}

func batchInsert[T dbmodels.CourseInfo | dbmodels.TimeInfo](data []T) {
	batchSize := 500
	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		if err := database.Client.Save(data[i:end]).Error; err != nil {
			logger.Error(err)
			return
		}
	}
}
