package main

import (
	"cengkeHelperDev/service"
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/router"
	"cengkeHelperDev/src/utils/logger"
)

func main() {
	//UpdateDB()
	service.GetTeachInfos()
	return
	logger.Info("start")

	if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
		logger.Error("run server error: ", err)
		return
	}
}
