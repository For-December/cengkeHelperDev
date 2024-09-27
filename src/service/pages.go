package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/logger"
)

// GetWrapperWithPage 分页查询封装
// M: Model type
// R: Result type
// 降序排序
func GetWrapperWithPage[
	M dbmodels.PostRecord,
	R dbmodels.PostRecord](
	page int, pageSize int,
	orderRule string,
	conditions ...interface{}) *models.DividePageWrapper[R] {

	if len(conditions) == 0 {
		conditions = append(conditions, "")
	}
	//logger.Info(conditions)
	// 这里拼接字符串会有sql注入的风险
	//queryStr := ""
	//for _, condition := range conditions {
	//	queryStr += " " + condition
	//}

	if page < 1 {
		page = 1
	}
	if pageSize <= 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	var count int64
	if err := database.Client.Model(new(M)).Where(conditions[0], conditions[1:]...).
		Count(&count).Error; err != nil {
		logger.Error(err)
		return nil
	}
	total := int(count)

	resData := make([]R, 0)
	if err := database.Client.Model(new(M)).Where(conditions[0], conditions[1:]...).
		Offset(offset).Limit(pageSize).
		Order(orderRule).
		Find(&resData).Error; err != nil {
		logger.Error(err)
		return nil
	}
	return &models.DividePageWrapper[R]{
		Page:     page,
		PageSize: pageSize,
		List:     resData,
		Total:    total,
	}

}
