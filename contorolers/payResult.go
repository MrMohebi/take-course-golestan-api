package contorolers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PayResult() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(http.StatusOK, "payResult.html", nil)
	}
}
