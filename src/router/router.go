package router

import (
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/calc"
	"cengkeHelperDev/src/utils/location"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var app *gin.Engine
var validHosts = []string{
	"localhost", "cengkehelper.top", "huster.pages.dev", "ursb.top", "蹭课小助手.top",
}

func Routers() *gin.Engine {

	app.POST("/teach-infos", func(c *gin.Context) {

		if !calc.IsTargetInArray(c.Request.Host, validHosts) {
			logger.WarningF("请求的主机不合法: %v, refer为: %v", c.Request.Host, c.Request.Referer())
			c.JSON(400, "bad request")
		} else {
			c.JSON(200, service.GetTeachInfos(true))
			//logger.Warning(GetTeachInfos(true))
			//c.JSON(200, GetTeachInfos(true))
		}

	})
	app.GET("/cur-time", func(c *gin.Context) {
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
	})
	app.POST("/register", func(c *gin.Context) {

	})
	return app

}

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	app = gin.Default()

	// 中间件，解决开发时的跨域问题
	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// 前端界面
	app.Use(static.Serve("/", static.LocalFile("dist", true)))

	// 静态文件的文件夹相对目录
	app.StaticFS("/dist", http.Dir("./dist"))
	// 单文件路径映射
	app.StaticFile("/favicon.ico", "./favicon.ico")
	app.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("dist1/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				_, err := c.Writer.WriteString("Not Found")
				if err != nil {
					logrus.Warning(err)
					return
				}
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			_, err = c.Writer.Write(content)
			if err != nil {
				return
			}
			c.Writer.Flush()
		}
	})
}
