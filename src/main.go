package main

import (
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/router"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
	"cengkeHelperDev/src/utils/web"
	"os"
)

func main() {

	file, err := os.OpenFile("dist/vite.svg", os.O_RDONLY, 0666)
	if err != nil {
		logger.Error(err)
	}

	niu, err := web.UploadToQiNiu(file)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(niu)
	return
	//UpdateDB()
	service.GetTeachInfos(false)
	//return
	logger.Info("start")

	if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
		logger.Error("run server error: ", err)
		return
	}
}
