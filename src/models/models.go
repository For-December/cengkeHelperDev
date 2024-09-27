package models

type CourseInfoRaw struct {
	ID         string `json:"id"`
	Years      string `json:"years"`     // 2021-2022
	Semester   string `json:"semester"`  // 秋季学期
	CourseNum  string `json:"courseNum"` // 2021-2022-1-1001
	CourseName string `json:"courseName"`
	Faculty    string `json:"faculty"`

	Credit string `json:"credit"` // 2.0

	CourseComplexion string `json:"courseComplexion"`
	CourseType       string `json:"courseType"`
	Grade            string `json:"grade"`
	Major            string `json:"major"`
	Teacher          string `json:"teacher"`
	TeacherTitle     string `json:"teacherTitle"`
	AddrSunday       string `json:"addrSunday"`
	AddrMonday       string `json:"addrMonday"`
	AddrTuesday      string `json:"addrTuesday"`
	AddrWednesday    string `json:"addrWednesday"`
	AddrThursday     string `json:"addrThursday"`
	AddrFriday       string `json:"addrFriday"`
	AddrSaturday     string `json:"addrSaturday"`
}
