package web

import (
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/constant/define"
	"cengkeHelperDev/src/utils/logger"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/objects"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"io"
	"mime/multipart"
	"time"
)

// MultiUploadToQiNiu 批量上传文件到七牛云
func MultiUploadToQiNiu(images []*multipart.FileHeader) ([]string, error) {

	// 并发上传图片
	errCh := make(chan error, len(images))

	// 这里初始化而不扩容切片，因为在循环中会对切片进行赋值，扩容会导致切片地址变化
	resImageURLs := make([]string, len(images)) // 保证图片的顺序

	for idx, image := range images {
		// 闭包捕获循环变量
		go func(index int, img *multipart.FileHeader) {
			// 打开文件
			file, err := img.Open()
			if err != nil {
				logger.Warning(err)
				errCh <- err
				return
			}
			defer func(file multipart.File) {
				err := file.Close()
				if err != nil {
					logger.Error(err)
				}
			}(file)

			imageStr, err := UploadToQiNiu(file)
			if err != nil {
				logger.Warning(err)
				errCh <- err
				return
			}
			resImageURLs[index] = imageStr
			errCh <- nil
		}(idx, image)
	}

	// 等待所有图片上传完成并检查错误
	for range images {
		if err := <-errCh; err != nil {
			return nil, err
		}
	}

	return resImageURLs, nil
}

// UploadToQiNiu 上传文件到七牛云
func UploadToQiNiu(reader io.Reader) (string, error) {
	accessKey := config.EnvCfg.QiNiuAccessKey
	secretKey := config.EnvCfg.QiNiuSecretKey
	bucket := define.QiNiuBucket

	mac := credentials.NewCredentials(accessKey, secretKey)

	key := fmt.Sprintf("img_%d", time.Now().Unix())

	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err := uploadManager.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: bucket,
		ObjectName: &key,
		FileName:   key,
	}, nil)
	if err != nil {
		logger.Error("上传失败，可能是key和secret配置有误：", err)
		return "", err
	}

	return define.QiNiuDomain + "/" + key, err

}

func DeleteFromQiNiu(key string) error {
	accessKey := config.EnvCfg.QiNiuAccessKey
	secretKey := config.EnvCfg.QiNiuSecretKey
	bucketName := define.QiNiuBucket

	mac := credentials.NewCredentials(accessKey, secretKey)
	objectsManager := objects.NewObjectsManager(&objects.ObjectsManagerOptions{
		Options: http_client.Options{Credentials: mac},
	})

	bucket := objectsManager.Bucket(bucketName)

	err := bucket.Object(key).Delete().Call(context.Background())
	if err != nil {
		return err
	}

	return nil
}
