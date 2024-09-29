package web

import (
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/constant/define"
	"cengkeHelperDev/src/utils/logger"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"io"
	"time"
)

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
		logger.Error("上传失败，可能是没配置key和secret：", err)
		return "", err
	}

	return define.QiNiuDomain + "/" + key, err

}
