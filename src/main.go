package main

import (
	"cengkeHelperDev/src/constant/config"
	"cengkeHelperDev/src/router"
	"cengkeHelperDev/src/utils/logger"
)

func main() {
	logger.Info("start")

	if err := router.Routers().Run(":" + config.EnvCfg.ServerPort); err != nil {
		logger.Error("run server error: ", err)
		return
	}
}
