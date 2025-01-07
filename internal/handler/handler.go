package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"test": "test",
	})
}
