package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary      获得用户列表
// @Description  获得用户列表
// @Tags         用户相关接口
// @Accept       json
// @Produce      json
// @Success      200  {string} 	welcome
// @Failure      400  {string} 	welcome
// @Failure      404  {string} 	welcome
// @Failure      500  {string} 	welcome
// @Router       /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
