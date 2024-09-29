package tree_hole

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostsGetAllHandler(c *gin.Context) {

	page, pageSize, searchValue := service.ParsePageParams(c)
	logger.Info(page, pageSize, searchValue)

	//res := service.GetPostsWithPage(page, pageSize,
	//	"title LIKE ?",
	//	"%"+searchValue+"%")
	//if res == nil {
	//	// 断言不会进入这里
	//	c.JSON(http.StatusBadRequest,
	//		models.NewBadResp("分页失败，请联系开发者"))
	//	return
	//}

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

func PostsCreateOneHandler(c *gin.Context) {
	var post dbmodels.PostRecord
	if err := c.ShouldBindJSON(&post); err != nil {
		logger.Warning(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	if err := service.CreatePost(&post); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("发帖失败！"))
		return
	}

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: post,
		Msg:  "success",
	})
}
