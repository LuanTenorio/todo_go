package router

import (
	"github.com/gin-gonic/gin"
)

func Inicialize() {
	r := gin.Default()
	inicializeRoutes(r)
	r.Run(":8080")
}
