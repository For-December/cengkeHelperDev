package tree_hole

import (
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func StarsUpdateOneHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	// 从token中获取用户id
	userIdStr := "1"
	userId, _ := strconv.Atoi(userIdStr)

	// 从请求中获取点赞状态
	type T struct {
		IsStar bool `json:"isStar"`
	}
	req := T{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	// 修改点赞状态
	if err := service.UpdateStar(
		uint32(userId),
		uint32(id),
		req.IsStar,
	); err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("点赞失败！"))
		return
	}

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

func StarsGetAllHandler(c *gin.Context) {

}
