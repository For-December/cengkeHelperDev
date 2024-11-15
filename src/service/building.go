package service

import (
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/generator"
	"cengkeHelperDev/src/utils/logger"
	"slices"
	"time"
)

type RespTeachInfo struct {
	Room         string `json:"room"`
	Faculty      string `json:"faculty"`
	CourseName   string `json:"courseName"`
	TeacherName  string `json:"teacherName"`
	TeacherTitle string `json:"teacherTitle"`
	CourseTime   string `json:"courseTime"`
	CourseType   string `json:"courseType"`
}
type MapTeachInfo struct {
	Classroom    string
	Faculty      string
	CourseName   string
	Teacher      string
	TeacherTitle string
	WeekAndTime  uint32
	Building     string

	CourseType string
}

// BuildingTeachInfos 每个学部各个教学楼的课程信息
type BuildingTeachInfos struct {
	Building string          `json:"building"`
	Infos    []RespTeachInfo `json:"infos"`
}

var RespTeachInfos = make([][]BuildingTeachInfos, 5)

func searchByArea(areaNum int) []MapTeachInfo {
	tempInfo := make([]MapTeachInfo, 0)
	if err := database.Client.
		Raw(`
SELECT * FROM time_infos ti 
         JOIN course_infos ci on ci.id = ti.course_info_id
         WHERE ti.day_of_week = ? AND ti.area = ? GROUP BY building,classroom`,
			time.Now().Weekday(), areaNum).
		Find(&tempInfo).Error; err != nil {
		logger.Error(err)
	}

	return tempInfo
}
func getInfos() [][]BuildingTeachInfos {
	//tempCourse := make([]dbmodels.CourseInfo, 0)
	for i := 0; i < 5; i++ {
		RespTeachInfos[i] = make([]BuildingTeachInfos, 0)
	}
	for i := 1; i <= 4; i++ {
		buildingMap := make(map[string][]RespTeachInfo)
		for _, info := range searchByArea(i) {
			weekNum, _, lessonNum := CurCourseTime()
			if !generator.IsWeekLessonMatch(weekNum, lessonNum, info.WeekAndTime) {
				continue
			}

			res := RespTeachInfo{
				Room:         info.Classroom,
				Faculty:      info.Faculty,
				CourseName:   info.CourseName,
				TeacherName:  info.Teacher,
				TeacherTitle: info.TeacherTitle,
				CourseTime:   generator.NearestToDisplay(lessonNum, info.WeekAndTime),
				CourseType:   info.CourseType,
			}
			//_, lesson := generator.Bin2WeekLesson(info.WeekAndTime)
			//logger.Info(res, lesson)
			buildingMap[info.Building] = append(buildingMap[info.Building], res)
		}
		for key, infos := range buildingMap {
			RespTeachInfos[i-1] = append(RespTeachInfos[i-1], BuildingTeachInfos{
				Building: key,
				Infos:    infos,
			})
		}

	}

	// 每个学部的教学楼按照课程数量排序
	for i := range RespTeachInfos {
		slices.SortFunc(RespTeachInfos[i], func(a, b BuildingTeachInfos) int {
			return len(b.Infos) - len(a.Infos)
		})
	}

	//for _, info := range searchByArea(1) {
	//	if !generator.IsWeekLessonMatch(2, 2, info.WeekAndTime) {
	//		continue
	//	}
	//	buildingMap := make(map[string][]RespTeachInfo)
	//	buildingMap[info.Building] = append(buildingMap[info.Building], RespTeachInfo{})
	//
	//}
	//logger.Warning(tempCourse)

	return RespTeachInfos
}
