package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Summary      得到信息p
// @Description  get string by ID
// @Tags         首页
// @Accept       json
// @Produce      json
// @Success      200  {string} 	welcome
// @Failure      400  {string} 	welcome
// @Failure      404  {string} 	welcome
// @Failure      500  {string} 	welcome
// @Router       /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "p",
	})
}
