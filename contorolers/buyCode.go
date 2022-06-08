package contorolers

import (
	"context"
	"github.com/MrMohebi/take-course-golestan-api.git/common"
	"github.com/MrMohebi/take-course-golestan-api.git/faces"
	"github.com/MrMohebi/take-course-golestan-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func BuyCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.LoginReq
		common.ValidBindForm(c, &reqBody)

		var user models.User

		token := common.RandStr(32)

		isJoined := true
		userNotFound := models.UsersCollection.FindOne(ctx, bson.M{"phone": reqBody.Phone}).Decode(&user)
		if userNotFound != nil {
			isJoined = false
		}

		if !isJoined {
			singUp(&reqBody, &token)
		}

		if isJoined && !login(&reqBody, &user, &token) {
			c.JSON(http.StatusUnauthorized, 401)
			return
		}

		//c.JSON(http.StatusOK, gin.H{
		//	"token": token,
		//})
	}
}
