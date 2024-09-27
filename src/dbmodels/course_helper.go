package dbmodels

type CourseInfo struct {
	ID         uint32 `gorm:"not null;primaryKey;autoIncrement" json:"id"`
	Years      string `gorm:"not null;type:varchar(255)" json:"years"`     // 2021-2022
	Semester   string `gorm:"not null;type:varchar(255)" json:"semester"`  // 秋季学期
	CourseNum  string `gorm:"not null;type:varchar(255)" json:"courseNum"` // 2021-2022-1-1001
	CourseName string `gorm:"not null;type:varchar(255)" json:"courseName"`
	Faculty    string `gorm:"not null;type:varchar(255)" json:"faculty"`

	Credit string `gorm:"not null;type:varchar(255)" json:"credit"` // 2.0

	CourseComplexion string `gorm:"not null;type:varchar(255)" json:"courseComplexion"`
	CourseType       string `gorm:"not null;type:varchar(255)" json:"courseType"`
	Grade            string `gorm:"not null;type:varchar(255)" json:"grade"`
	Major            string `gorm:"not null;type:varchar(255)" json:"major"`
	Teacher          string `gorm:"not null;type:varchar(255)" json:"teacher"`
	TeacherTitle     string `gorm:"not null;type:varchar(255)" json:"teacherTitle"`
}

type TimeInfo struct {
	ID           uint32 `gorm:"not null;primaryKey;autoIncrement" json:"id"`
	CourseInfoId uint32 `gorm:"not null" json:"courseInfo"`

	WeekAndTime uint32 `gorm:"not null" json:"weekAndTime"`

	DayOfWeek uint8 `gorm:"not null" json:"dayOfWeek"` // 0-6

	Area      uint8  `gorm:"not null" json:"area"` // 1-4
	Building  string `gorm:"not null;type:varchar(255)" json:"building"`
	Classroom string `gorm:"not null;type:varchar(255)" json:"classroom"`
}
