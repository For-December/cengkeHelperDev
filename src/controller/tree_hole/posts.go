package tree_hole

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostsGetAllHandler(c *gin.Context) {
	page := 1
	pageSize := 10
	searchValue := ""
	if value, exists := c.GetQuery("page"); exists {
		page, _ = strconv.Atoi(value)
	}

	if value, exists := c.GetQuery("pageSize"); exists {
		pageSize, _ = strconv.Atoi(value)

	}

	if value, exists := c.GetQuery("searchValue"); exists {
		searchValue = value
	}

	//res := service.GetPostsWithPage(page, pageSize,
	//	"title LIKE ?",
	//	"%"+searchValue+"%")
	//if res == nil {
	//	// 断言不会进入这里
	//	c.JSON(http.StatusBadRequest,
	//		models.NewBadResp("分页失败，请联系开发者"))
	//	return
	//}
	logger.Info(page, pageSize, searchValue)

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: struct {
			Page     int                   `json:"page"`
			PageSize int                   `json:"pageSize"`
			List     []dbmodels.PostRecord `json:"list"`
			Total    int                   `json:"total"`
		}{
			Page:     1,
			PageSize: 10,
			List: []dbmodels.PostRecord{
				{
					BaseModel:    dbmodels.BaseModel{},
					AuthorID:     1,
					AuthorName:   "jack",
					CommentCount: 1,
					UpvoteCount:  2,
					Title:        "测试",
					ContentJson: "[{\"type\":\"text\",\"text\":\"hello\"}," +
						"{\"type\":\"image\"," +
						"\"url\":\"https://tse4-mm.cn.bing.net/th/id/OIP-C.YKoZzgmubNBxQ8j-mmoTKAHaEK?rs=1&pid=ImgDetMain\"}]",
				},
			},
			Total: 1,
		},
		Msg: "success",
	})
}
