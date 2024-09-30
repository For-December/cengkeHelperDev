package service

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/storage/database"
	"cengkeHelperDev/src/utils/logger"
	"cengkeHelperDev/src/utils/web"
	"mime/multipart"
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
	}

	return post
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
