package web

import (
	"cengkeHelperDev/src/utils/logger"
	"mime/multipart"
	"os"
	"testing"
)

func TestDeleteFromQiNiu(t *testing.T) {
	err := DeleteFromQiNiu("img_1728355351")
	if err != nil {
		logger.Error(err)
		return
	}
}

func TestUploadToQiNiu(t *testing.T) {
	file, err := os.OpenFile("dist/vite.svg", os.O_RDONLY, 0666)
	if err != nil {
		logger.Error(err)
	}

	niu, err := UploadToQiNiu(file)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(niu)
	return
}

func TestMultiUploadToQiNiu(t *testing.T) {

	niu, err := MultiUploadToQiNiu(make([]*multipart.FileHeader, 3))
	if err != nil {
		logger.Error(err)
		return
	}
	for _, s := range niu {
		logger.Info(s)
	}

}
