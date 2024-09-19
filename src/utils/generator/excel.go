package generator

import (
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/utils/calc"
	"cengkeHelperDev/src/utils/logger"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

var usefulInfoCols = []string{
	"学年", "学期", "课程号", "课程名称", "开课学院", "学分", "课程性质", "开课类型", "年级", "专业",
	"成绩录入老师", "职称", "周日", "周一", "周二", "周三", "周四", "周五", "周六",
}

var infoExcel2Struct = map[string]string{
	"学年":     "Years",
	"学期":     "Semester",
	"课程号":    "CourseNum",
	"课程名称":   "CourseName",
	"开课学院":   "Faculty",
	"学分":     "Credit",
	"课程性质":   "CourseComplexion",
	"开课类型":   "CourseType",
	"年级":     "Grade",
	"专业":     "Major",
	"成绩录入老师": "Teacher",
	"职称":     "TeacherTitle",
	"周日":     "AddrSunday",
	"周一":     "AddrMonday",
	"周二":     "AddrTuesday",
	"周三":     "AddrWednesday",
	"周四":     "AddrThursday",
	"周五":     "AddrFriday",
	"周六":     "AddrSaturday",
}

func ReadTeachInfo() (resInfos []models.CourseInfoRaw, err error) {
	return readExcel[models.CourseInfoRaw](
		"resources/2024-2025学年第一学期全校课表.xlsx",
		usefulInfoCols, infoExcel2Struct)
}

func readExcel[T models.CourseInfoRaw](
	filePath string,
	usefulCols []string,
	excel2Struct map[string]string,
) (resInfos []T, err error) {

	var f *excelize.File

	// 读取数据
	if f, err = excelize.OpenFile(filePath); err != nil {
		logger.Error(err)
		return
	}
	defer func(f *excelize.File) {
		if err := f.Close(); err != nil {
			logger.Error(err)
		}
	}(f)

	// 获取 data 上所有单元格(第一个Sheet)
	// 突然发现有的是 data 有的是 sheet1，所以这里直接获取第一个sheet
	var rows [][]string
	sheets := f.GetSheetList()
	if len(sheets) < 1 {
		err = errors.New("sheet数量小于1")
		return
	}
	// 直接读第一个sheet
	if rows, err = f.GetRows(sheets[0]); err != nil {
		fmt.Println(err)
		return
	}

	// 获取有用行的索引
	usefulIndex := make([]int, 0)
	for index, titleCel := range rows[0] {
		if calc.IsTargetInArray[string](titleCel, usefulCols) {
			usefulIndex = append(usefulIndex, index)
		}
	}

	logger.Warning(usefulCols)
	logger.Warning(usefulIndex)

	// 保存数据
	resInfos = make([]T, 0)

	// 查询数据
	for _, row := range rows[1:] {

		// 反射创建结构体来记录每一行的数据
		// 直接 TypeOf(T) 无法通过编译，因此这里用 new(T) 来获取类型
		// 由于指定了多个泛型、无法确定结构体所占空间大小?
		// 参见 -> https://go.dev/doc/tutorial/generics
		resInfo := reflect.New(reflect.TypeOf(*new(T))).Elem()

		for j, colCell := range row {
			if !calc.IsTargetInArray[int](j, usefulIndex) {
				continue
			}

			//if len(colCell) == 0 {
			//	err = errors.New("表格存在无效数据")
			//	return
			//}

			// 通过反射将对应行的数据保存到结构体中
			field := resInfo.FieldByName(excel2Struct[rows[0][j]])

			switch field.Type() {
			case reflect.TypeOf(uint32(0)):
				num, _ := strconv.Atoi(colCell)
				field.SetUint(uint64(num))
			case reflect.TypeOf(""):
				field.SetString(colCell)
			case reflect.TypeOf(float32(0)):
				floatData, _ := strconv.ParseFloat(colCell, 32)
				field.SetFloat(floatData)
			default:
				logger.Error("未知类型" + field.Type().String())
			}

		}

		// 打印读取到的数据
		//logger.Info(resInfo.Interface())

		// 只保存这一行需要的数据，并添加到结果数组
		resInfos = append(resInfos, resInfo.Interface().(T))

	}
	return
}
