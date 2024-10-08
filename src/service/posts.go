package service

import (
	"cengkeHelperDev/src/constant/define"
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/logger"
	"cengkeHelperDev/src/utils/web"
	"encoding/json"
	"errors"
	"mime/multipart"
	"strings"
	"time"
)

func GetPostsWithPage(page int,
	pageSize int,
	condition ...interface{}) *models.DividePageWrapper[dbmodels.PostRecord] {
	return GetWrapperWithPage[dbmodels.PostRecord, dbmodels.PostRecord](
		page, pageSize, "latest_replied_at desc", condition...)
}

func GetPostById(id uint32) *dbmodels.PostRecord {
	post := &dbmodels.PostRecord{}
	if err := database.Client.First(post, id).Error; err != nil {
		logger.Warning(err)
		return nil
	}
	if post.ID == 0 {
		logger.WarningF("post %d not found", id)
		return nil
	}

	return post
}

func DeletePostById(id uint32) error {
	post := &dbmodels.PostRecord{}
	if err := database.Client.First(post, id).Error; err != nil {
		logger.Warning(err)
		return err
	}
	if post.ID == 0 {
		logger.WarningF("post %d not found", id)
		return errors.New("贴子不存在！")
	}

	// 先从数据库中删除贴子
	if err := database.Client.Model(&dbmodels.PostRecord{}).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		logger.Warning(err)
		return err
	}

	// 删除对象存储中的图片资源
	postMetas := make([]models.PostMeta, 0)
	err := json.Unmarshal([]byte(post.ContentJson), &postMetas)
	if err != nil {
		logger.Warning(err)
		return err
	}

	// 删除贴子的所有图片
	for _, meta := range postMetas {
		if meta.Url != "" {
			after, found := strings.CutPrefix(meta.Url, define.QiNiuDomain+"/")
			if !found {
				logger.Warning("不符合定义的url: ", meta.Url)
				continue
			}

			if err := web.DeleteFromQiNiu(after); err != nil {
				logger.Warning(err)
				return err
			}
		}
	}

	// 删除帖子

	return nil
}

func SavePost(post *dbmodels.PostRecord) error {
	return database.Client.Model(&dbmodels.PostRecord{}).Save(post).Error
}

func CreatePost(authorId uint32, authorName string, text string, images []*multipart.FileHeader) error {
	// 使用上下文来控制操作的超时或取消
	//ctx := context.Background()

	postBuilder := models.PostMetaBuilder{}
	postBuilder.BuildText(text)

	resImageURLs, err := web.MultiUploadToQiNiu(images)
	if err != nil {
		logger.Warning(err)
		return err
	}
	postBuilder.BuildImages(resImageURLs)

	post := &dbmodels.PostRecord{
		BaseModel:       dbmodels.BaseModel{},
		AuthorId:        authorId,
		AuthorName:      authorName,
		CommentCount:    0,
		UpvoteCount:     0,
		Title:           "",
		ContentJson:     postBuilder.BuildJson(),
		LatestRepliedAt: time.Now(),
	}

	return database.Client.Create(post).Error
}
