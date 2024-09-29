package web

import (
	"cengkeHelperDev/src/utils/logger"
	"os"
	"testing"
)

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
