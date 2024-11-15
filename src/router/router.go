package router

import (
	"cengkeHelperDev/src/controller/course_helper"
	"cengkeHelperDev/src/controller/tree_hole"
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

func Routers() *gin.Engine {

	app.POST("/teach-infos", course_helper.TeachInfoHandler)
	app.GET("/cur-time", course_helper.CurTimeHandler)

	// 树洞api v1
	v1 := app.Group("/api/v1")
	{
		v1.GET("posts", tree_hole.PostsGetAllHandler)
		v1.POST("posts", tree_hole.PostsCreateOneHandler)
		v1.GET("posts/:id", tree_hole.PostsGetOneHandler)
		v1.DELETE("posts/:id", tree_hole.PostsDeleteOneHandler)

		//v1.POST("posts/:id/upvote", tree_hole.PostsUpvoteHandler)

		v1.GET("posts/:id/comments", tree_hole.PostsGetCommentsHandler)
		v1.POST("posts/:id/comments", tree_hole.CommentsCreateOneHandler)

		// 修改某个贴子的点赞状态
		v1.POST("posts/:id/stars", tree_hole.StarsUpdateOneHandler)

		// 删除评论
		v1.DELETE("comments/:id", tree_hole.CommentsDeleteOneHandler)

		// 个人点赞信息
		v1.GET("profile/stars", tree_hole.StarsGetAllHandler)

	}
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
