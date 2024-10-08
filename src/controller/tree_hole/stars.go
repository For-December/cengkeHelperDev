package tree_hole

import (
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func StarsCreateOneHandler(c *gin.Context) {
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

	// 保存点赞
	if err := service.SaveStar(
		uint32(userId),
		uint32(id),
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
