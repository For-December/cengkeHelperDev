package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/logger"
	"errors"
	"gorm.io/gorm"
	"time"
)

func SaveStar(userId, postId uint32) error {
	if err := database.Client.Transaction(func(tx *gorm.DB) error {
		// 贴子点赞数+1
		if rowsAffected := tx.Model(&dbmodels.PostRecord{}).
			Where("id = ?", postId).
			UpdateColumn("upvote_count", gorm.Expr("upvote_count + 1")).
			RowsAffected; rowsAffected != 1 {
			return errors.New("已点赞！")
		}

		// 保存点赞信息
		star := &dbmodels.StarRecord{
			PostId:    postId,
			UserId:    userId,
			CreatedAt: time.Now(),
		}

		if err := tx.Save(star).Error; err != nil {
			return err
		}

		return nil

	}); err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}
