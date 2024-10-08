package service

import (
	"cengkeHelperDev/src/constant/define"
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/logger"
	"errors"
	"gorm.io/gorm"
	"time"
)

func UpdateStar(userId, postId uint32, isStar bool) error {
	if err := database.Client.Transaction(func(tx *gorm.DB) error {

		// 贴子点赞数+1
		sqlExpr := gorm.Expr("upvote_count + 1")

		// 如果取消点赞，贴子点赞数-1
		if !isStar {
			sqlExpr = gorm.Expr("upvote_count - 1")
		}

		if rowsAffected := tx.Model(&dbmodels.PostRecord{}).
			Where("id = ?", postId).
			UpdateColumn("upvote_count", sqlExpr).
			RowsAffected; rowsAffected != 1 {
			return errors.New("贴子不存在！")
		}

		// 保存或删除点赞信息
		var count int64
		if err := tx.Model(&dbmodels.StarRecord{}).
			Where("post_id = ? AND user_id = ?", postId, userId).
			Count(&count).Error; err != nil {
			logger.Warning(err)
			return err
		}

		if isStar {
			if count != 0 {
				logger.Warning("已点过赞，状态未改变")
				return define.RecoverableError
			}

			star := &dbmodels.StarRecord{
				PostId:    postId,
				UserId:    userId,
				CreatedAt: time.Now(),
			}

			if err := tx.Create(star).Error; err != nil {
				return err
			}
		} else {
			if count == 0 {
				logger.Warning("未点赞，状态未改变")
				return define.RecoverableError
			}

			if err := tx.Model(&dbmodels.StarRecord{}).
				Where("post_id = ? AND user_id = ?", postId, userId).
				Delete(nil).Error; err != nil {
				return err
			}
		}

		return nil

	}); err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}

func GetStarsList(userId uint32) []dbmodels.StarRecord {

	starRecords := make([]dbmodels.StarRecord, 0)
	if err := database.Client.
		Model(&dbmodels.StarRecord{}).
		Where("user_id = ?", userId).
		Find(&starRecords).Error; err != nil {
		logger.Warning(err)
	}

	return starRecords
}
