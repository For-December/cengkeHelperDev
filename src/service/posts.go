package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
)

func GetPostsWithPage(page int,
	pageSize int,
	condition ...interface{}) *models.DividePageWrapper[dbmodels.PostRecord] {
	return GetWrapperWithPage[dbmodels.PostRecord, dbmodels.PostRecord](
		page, pageSize, "latest_replied_at desc", condition...)
}
