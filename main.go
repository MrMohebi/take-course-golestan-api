package main

import (
	"github.com/MrMohebi/take-course-golestan-api.git/common"
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"github.com/MrMohebi/take-course-golestan-api.git/router"
	"github.com/gin-gonic/gin"
)

// nodemon --exec go run main.go --signal SIGTERM

func main() {
	configs.Setup()

	server := gin.Default()

	router.Routs(server)

	err := server.Run("localhost:80")
	common.IsErr(err, "Err in starting server")
}
