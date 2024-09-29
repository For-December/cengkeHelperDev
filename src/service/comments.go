package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/storage/database"
	"errors"
	"gorm.io/gorm"
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

func SaveComment(authorId, postId uint32, authorName, content string) error {

	if err := database.Client.Transaction(func(tx *gorm.DB) error {
		// 获取当前帖子的信息
		post := dbmodels.PostRecord{}
		if err := tx.Model(&dbmodels.PostRecord{}).
			Where("id = ?", postId).
			Find(&post).Error; err != nil {
			return err
		}

		if post.ID == 0 {
			return errors.New("帖子不存在！")
		}

		// 保存评论
		comment := &dbmodels.CommentRecord{
			PostId:     postId,
			AuthorId:   authorId,
			AuthorName: authorName,
			Content:    content,
			FloorNum:   uint32(post.CommentCount) + 1,
		}

		if err := tx.Save(comment).Error; err != nil {
			return err
		}

		// 更新帖子的评论数
		if err := tx.Model(&dbmodels.PostRecord{}).
			Where("id = ?", postId).
			Update("comment_count", post.CommentCount+1).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil

}
