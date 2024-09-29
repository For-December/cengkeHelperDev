package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/storage/database"
)

func GetPostsWithPage(page int,
	pageSize int,
	condition ...interface{}) *models.DividePageWrapper[dbmodels.PostRecord] {
	return GetWrapperWithPage[dbmodels.PostRecord, dbmodels.PostRecord](
		page, pageSize, "latest_replied_at desc", condition...)
}

func SavePost(post *dbmodels.PostRecord) error {
	return database.Client.Model(&dbmodels.PostRecord{}).Save(post).Error
}

func CreatePost(post *dbmodels.PostRecord) error {
	post.ID = 0
	return database.Client.Model(&dbmodels.PostRecord{}).Create(post).Error
}
