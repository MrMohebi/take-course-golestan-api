package router

import (
	"github.com/MrMohebi/take-course-golestan-api.git/contorolers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routs(r *gin.Engine) {
	AssetsRoute(r)
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", contorolers.Index())
	r.GET("/docs", contorolers.Docs())
	r.GET("/payResult", contorolers.PayResult())
	r.GET("/aboutUs", func(context *gin.Context) { context.HTML(http.StatusOK, "aboutUs.html", nil) })

	r.POST("/login", contorolers.Login())
	r.POST("/buyCode", contorolers.BuyCode())
	r.POST("/payVerify", contorolers.PayVerify())
	r.POST("/isActiveCodeValid", contorolers.IsActiveCodeValid())
}
