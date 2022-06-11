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

func IsActiveCodeValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.IsActiveCodeValidReq
		isValidateReq := common.ValidBindForm(c, &reqBody)
		if !isValidateReq {
			return
		}

		var activeCode models.ActiveCode

		notFound := models.ActiveCodesCollection.FindOne(ctx, bson.M{"code": reqBody.Code}).Decode(&activeCode)
		if notFound != nil {
			c.JSON(http.StatusOK, faces.IsActiveCodeValidRes{IsValid: false, DaysLeft: 0})
			return
		}

		daysLeft := activeCodeDaysLeft(activeCode.CreatedAt)
		if daysLeft < 1 {
			c.JSON(http.StatusOK, faces.IsActiveCodeValidRes{IsValid: false, DaysLeft: 0})
			return
		}

		c.JSON(http.StatusOK, faces.IsActiveCodeValidRes{IsValid: true, DaysLeft: daysLeft})
	}
}
