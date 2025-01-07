package router

import (
	"github.com/LuanTenorio/todo_go/internal/handler"
	"github.com/gin-gonic/gin"
)

func inicializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("test", handler.TestHandler)
}
