package main

import (
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/router"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
)

func main() {
	//UpdateDB()
	service.GetTeachInfos(false)
	//return
	logger.Info("start")

	if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
		logger.Error("run server error: ", err)
		return
	}
}
