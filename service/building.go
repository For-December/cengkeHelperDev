package service

import (
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/generator"
	"cengkeHelperDev/src/utils/logger"
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

	CourseType string
}

// BuildingTeachInfos 每个学部各个教学楼的课程信息
type BuildingTeachInfos struct {
	Building string          `json:"building"`
	Infos    []RespTeachInfo `json:"infos"`
}

var respTeachInfos = make([][]BuildingTeachInfos, 5)

func searchByArea(areaNum int) []MapTeachInfo {
	tempInfo := make([]MapTeachInfo, 0)
	if err := database.Client.
		Raw(`
SELECT * FROM time_infos ti 
         JOIN course_infos ci on ci.id = ti.course_info_id
         WHERE ti.day_of_week = ? AND ti.area = ?`,
			time.Now().Weekday(), areaNum).
		Find(&tempInfo).Error; err != nil {
		logger.Error(err)
	}

	return tempInfo
}
func GetTeachInfos() [][]BuildingTeachInfos {
	//tempCourse := make([]dbmodels.CourseInfo, 0)

	for _, info := range searchByArea(1) {
		if !generator.IsWeekLessonMatch(2, 2, info.WeekAndTime) {
			continue
		}
		res := RespTeachInfo{
			Room:         info.Classroom,
			Faculty:      info.Faculty,
			CourseName:   info.CourseName,
			TeacherName:  info.Teacher,
			TeacherTitle: info.TeacherTitle,
			CourseTime:   "",
			CourseType:   info.CourseType,
		}
		logger.Info(res)

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

	return respTeachInfos
}
