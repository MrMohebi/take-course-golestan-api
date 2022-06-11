package router

import (
	"github.com/MrMohebi/take-course-golestan-api.git/contorolers"
	"github.com/gin-gonic/gin"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", contorolers.Index())
	r.GET("/docs", contorolers.Docs())

	r.POST("/login", contorolers.Login())
	r.POST("/buyCode", contorolers.BuyCode())
	r.POST("/payVerify", contorolers.PayVerify())
	r.POST("/isActiveCodeValid", contorolers.IsActiveCodeValid())
}
