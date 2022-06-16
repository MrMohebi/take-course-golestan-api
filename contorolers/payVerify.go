package contorolers

import (
	"context"
	"github.com/MrMohebi/take-course-golestan-api.git/common"
	"github.com/MrMohebi/take-course-golestan-api.git/common/idpay"
	"github.com/MrMohebi/take-course-golestan-api.git/faces"
	"github.com/MrMohebi/take-course-golestan-api.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func PayVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentTime := time.Now().Unix()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reqBody faces.PayVerifyReq
		if !common.ValidBindForm(c, &reqBody) {
			return
		}

		var ipg = ""
		if reqBody.IdpayId != "" {
			ipg = "idpay"
		}

		isVerified := false
		if ipg == "idpay" {
			if reqBody.Status == 10 {
				isVerified = idpayVerify(&reqBody)
			}
		}

		if isVerified {
			var payment models.Payment
			paymentNotFound := models.PaymentsCollection.FindOne(ctx, bson.M{"orderId": reqBody.OrderId}).Decode(&payment)
			if paymentNotFound != nil {
				common.IsErr(paymentNotFound)
				c.JSON(http.StatusInternalServerError, gin.H{})
				location := url.URL{Path: "/payResult"}
				c.Redirect(http.StatusFound, location.RequestURI())
				return
			}
			_, err := models.ActiveCodesCollection.InsertOne(ctx, models.ActiveCode{
				Id:        primitive.NewObjectID(),
				Code:      payment.OrderId,
				Phone:     payment.PayerPhone,
				CreatedAt: currentTime,
			})
			if err == nil {
				intCode, _ := strconv.Atoi(payment.OrderId)
				if sendActiveCode(intCode, payment.PayerPhone) {
					q := url.Values{}
					q.Set("activeCode", payment.OrderId)
					q.Set("sentSMS", "1")
					location := url.URL{Path: "/payResult", RawQuery: q.Encode()}
					c.Redirect(http.StatusFound, location.RequestURI())
					return
				}
			}
		}
		location := url.URL{Path: "/payResult"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

func idpayVerify(reqBody *faces.PayVerifyReq) bool {
	currentTime := time.Now().Unix()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var payment models.Payment
	paymentNotFound := models.PaymentsCollection.FindOne(ctx, bson.M{"idpayId": reqBody.IdpayId}).Decode(&payment)
	if paymentNotFound != nil {
		common.IsErr(paymentNotFound)
		return false
	}
	// is valid
	if !((reqBody.Amount == payment.Amount) && (reqBody.OrderId == payment.OrderId)) {
		return false
	}

	isTestPay, _ := strconv.ParseBool(os.Getenv("IS_PAYMENT_DEV"))
	paymentClient := idpay.NewClient(os.Getenv("IDPAY_KEY"), isTestPay)
	verifyRes := paymentClient.Verify(payment.IdpayId, payment.OrderId)

	if !(verifyRes.ReqStatus.Success) {
		println(verifyRes.ErrorMessage)
		println(verifyRes.Status)
		return false
	}

	// update fields of payment
	payment.VerifiedAt = currentTime
	payment.PaidAt = verifyRes.Payment.Date
	payment.CardNumber = verifyRes.Payment.CardNo
	payment.CardNumberHash = verifyRes.Payment.HashedCardNo
	payment.IdpayTrackingId = verifyRes.TrackId

	_, err := models.PaymentsCollection.UpdateOne(ctx, bson.M{"idpayId": reqBody.IdpayId}, bson.D{{"$set", payment}})

	return err == nil
}
