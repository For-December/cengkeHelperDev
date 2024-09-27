package tree_hole

import (
	"cengkeHelperDev/src/dbmodels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostsGetAllHandler(c *gin.Context) {

	c.JSON(http.StatusOK, []dbmodels.PostRecord{
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
	})
}
