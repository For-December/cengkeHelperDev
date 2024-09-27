package course_helper

import (
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/calc"
	"cengkeHelperDev/src/utils/location"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
)

var validHosts = []string{
	"localhost", "cengkehelper.top", "huster.pages.dev", "ursb.top", "蹭课小助手.top",
}

func TeachInfoHandler(c *gin.Context) {
	if !calc.IsTargetInArray(c.Request.Host, validHosts) {
		logger.WarningF("请求的主机不合法: %v, refer为: %v", c.Request.Host, c.Request.Referer())
		c.JSON(400, "bad request")
	} else {
		c.JSON(200, service.GetTeachInfos(true))
		//logger.Warning(GetTeachInfos(true))
		//c.JSON(200, GetTeachInfos(true))
	}
}

func CurTimeHandler(c *gin.Context) {
	logger.InfoF("ip => %v 「%v」 ==> website", c.ClientIP(),
		location.IpToLocation(c.ClientIP()))

	weekNum, weekday, lessonNum := service.CurCourseTime()

	valid := true
	if !service.ValidCache() {
		logger.Warning("缓存失效")
		service.GetTeachInfos(false)
		service.FreshCacheFlag()
		valid = false
	}

	c.JSON(200, gin.H{
		"isAdjust":  false,
		"weekNum":   weekNum,
		"weekday":   weekday,
		"lessonNum": lessonNum,
		"valid":     valid,
	})
}
