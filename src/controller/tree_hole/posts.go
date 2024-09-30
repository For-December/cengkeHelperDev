package tree_hole

import (
	"cengkeHelperDev/src/dbmodels"
	"cengkeHelperDev/src/models"
	"cengkeHelperDev/src/service"
	"cengkeHelperDev/src/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostsGetAllHandler(c *gin.Context) {

	page, pageSize, searchValue := service.ParsePageParams(c)
	logger.Info(page, pageSize, searchValue)

	res := service.GetPostsWithPage(page, pageSize,
		"title LIKE ?",
		"%"+searchValue+"%")
	if res == nil {
		// 断言不会进入这里
		c.JSON(http.StatusBadRequest,
			models.NewBadResp("分页失败，请联系开发者"))
		return
	}
	res.List = append(res.List, dbmodels.PostRecord{
		BaseModel:    dbmodels.BaseModel{},
		AuthorId:     1,
		AuthorName:   "jack",
		CommentCount: 1,
		UpvoteCount:  2,
		Title:        "测试",
		ContentJson: "[{\"type\":\"text\",\"text\":\"hello\"}," +
			"{\"type\":\"image\"," +
			"\"url\":\"https://tse4-mm.cn.bing.net/th/id/OIP-C.YKoZzgmubNBxQ8j-mmoTKAHaEK?rs=1&pid=ImgDetMain\"}]",
	})

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: res,
		Msg:  "success",
	})
}

func PostsCreateOneHandler(c *gin.Context) {
	authorIdStr, _ := c.GetPostForm("authorId")
	authorName := c.PostForm("authorName")

	text := c.PostForm("text")
	form, _ := c.MultipartForm()
	images := form.File["images"]

	authorId, _ := strconv.Atoi(authorIdStr)

	if authorId == 0 || authorName == "" || text == "" {
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	if len(images) > 5 {
		c.JSON(http.StatusBadRequest, models.NewBadResp("图片数量超过5张！"))
		return
	}

	if err := service.CreatePost(uint32(authorId), authorName, text, images); err != nil {
		logger.Error(err)
		c.JSON(http.StatusBadRequest, models.NewBadResp("发帖失败！"))
		return
	}

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: nil,
		Msg:  "success",
	})
}

func PostsGetOneHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	post := service.GetPostById(uint32(id))
	if post == nil {
		c.JSON(http.StatusBadRequest, models.NewBadResp("帖子不存在！"))
		return
	}

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: post,
		Msg:  "success",
	})

}

func PostsGetCommentsHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewBadResp("参数错误！"))
		return
	}

	page, pageSize, _ := service.ParsePageParams(c)

	res := service.GetCommentsByPostIdWithPage(page, pageSize, uint32(id))
	if res == nil {
		c.JSON(http.StatusBadRequest, models.NewBadResp("分页失败，请联系开发者"))
		return
	}

	c.JSON(http.StatusOK, models.RespData{
		Code: 200,
		Data: res,
		Msg:  "success",
	})
}
