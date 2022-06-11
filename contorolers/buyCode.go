package contorolers

import (
	"context"
	"github.com/MrMohebi/take-course-golestan-api.git/common"
	"github.com/MrMohebi/take-course-golestan-api.git/common/idpay"
	"github.com/MrMohebi/take-course-golestan-api.git/faces"
	"github.com/MrMohebi/take-course-golestan-api.git/models"
	ghasedak "github.com/ghasedakapi/ghasedak-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

func BuyCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentTime := time.Now().Unix()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.BuyCodeReq
		isValidateReq := common.ValidBindForm(c, &reqBody)
		if !isValidateReq {
			return
		}

		var activeCode models.ActiveCode

		hasActiveCode := true
		userNotFound := models.ActiveCodesCollection.FindOne(ctx, bson.M{"phone": reqBody.Phone}).Decode(&activeCode)
		if userNotFound != nil {
			hasActiveCode = false
		}

		// send previous active code
		if hasActiveCode && (activeCodeDaysLeft(activeCode.CreatedAt) > 0) {
			code, err := strconv.Atoi(activeCode.Code)
			common.IsErr(err)
			if sendActiveCode(code, reqBody.Phone) {
				c.JSON(http.StatusOK, faces.BuyCodeRes{
					HasCode:  true,
					DaysLeft: activeCodeDaysLeft(activeCode.CreatedAt),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			return
		}

		// create payment link
		paymentInfo := models.Payment{
			Id:          primitive.NewObjectID(),
			Amount:      200000,
			CallbackUrl: "https://gtc.m3m.dev/payVerify",
			OrderId:     strconv.Itoa(common.RandNumber(11111111, 99999999)),
			PayerPhone:  reqBody.Phone,
			PayerEmail:  reqBody.Email,
			CreatedAt:   currentTime,
		}
		paymentClient := idpay.NewClient(os.Getenv("IDPAY_KEY"), false)
		createRes := paymentClient.CreatePayment(
			paymentInfo.OrderId,
			paymentInfo.Amount,
			paymentInfo.CallbackUrl,
			"",
			paymentInfo.PayerPhone,
			paymentInfo.PayerEmail,
			"",
		)

		if createRes.ReqStatus.Success {
			paymentInfo.Link = createRes.Link
			paymentInfo.IdpayId = createRes.Id
		}

		_, err := models.PaymentsCollection.InsertOne(ctx, paymentInfo)
		common.IsErr(err)

		if err == nil {
			c.JSON(http.StatusOK, faces.BuyCodeRes{
				HasCode: false,
				PayLink: paymentInfo.Link,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

func sendActiveCode(code int, phone string) bool {
	ghasedakClient := ghasedak.NewClient(os.Getenv("GHASEDAK_KEY"), "")
	r := ghasedakClient.SendOTP(phone, "gtcActiveCode", code)

	return r.Success
}

func activeCodeDaysLeft(createdAt int64) int {
	// 7 days
	deadline := 7
	currentTime := time.Now().Unix()
	passedSeconds := currentTime - createdAt
	passedDays := int(math.Floor(float64(passedSeconds / (24 * 60 * 60))))
	if leftDays := deadline - passedDays; leftDays > 0 {
		return leftDays
	}
	return 0
}
