package tree_hole

import (
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentsCreateOneHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	// 从token中获取用户id和用户名
	authorIdStr := "1"
	authorName := "芝士雪豹"
	authorId, _ := strconv.Atoi(authorIdStr)

	// 从请求中获取评论内容
	type T struct {
		Content string `json:"content"`
	}
	req := T{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	if req.Content == "" {
		c.JSON(http.StatusBadRequest, models.NewBadResp("评论内容不能为空！"))
		return
	}

	// 保存评论
	if err := service.SaveComment(
		uint32(authorId), uint32(id),
		authorName, req.Content); err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("评论失败！"))
		return
	}

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}
