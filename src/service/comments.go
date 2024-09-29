package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
)

// 内部使用
func getCommentsWithPage(page int,
	pageSize int,
	condition ...interface{}) *models.DividePageWrapper[dbmodels.CommentRecord] {
	return GetWrapperWithPage[dbmodels.CommentRecord, dbmodels.CommentRecord](
		page, pageSize, "created_at desc", condition...)
}

func GetCommentsByPostIdWithPage(
	page int,
	pageSize int,
	postId uint32) *models.DividePageWrapper[dbmodels.CommentRecord] {
	return getCommentsWithPage(page, pageSize, "post_id = ?", postId)
}
