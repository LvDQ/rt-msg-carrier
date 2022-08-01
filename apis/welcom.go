package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /
// Hello world godoc
// @Summary hello world example
// @Schemes
// @Description do hello-world
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /helloworld [get]
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "hello-world",
	})
}
