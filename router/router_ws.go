package router

import (
	ws_handler "rt-msg-carrier/apis/http/ws"

	"github.com/gin-gonic/gin"
)

// SetupRouter index
func SetupWsRouter(r *gin.Engine) *gin.Engine {
	ws := r.Group("/rt-msg-carrier/v1/ws")
	{
		ws.GET("", ws_handler.NewWsConn)
		ws.POST("/", ws_handler.PostWsMessageHandler)
	}

	return r
}
