package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	UserID          primitive.ObjectID `json:"username"`
	OrderId         string             `json:"orderId"`
	IdpayId         string             `json:"idpayId"`
	IdpayTrackingId string             `json:"idpayTrackingId"`
	PayerName       string             `json:"payerName"`
	PayerPhone      string             `json:"payerPhone"`
	PayerEmail      string             `json:"payerEmail"`
	CardNumber      string             `json:"cardNumber"`
	CardNumberHash  string             `json:"cardNumberHash"`
	Link            string             `json:"link"`
	CallbackUrl     string             `json:"callbackUrl"`
	Amount          int16              `json:"amount"`
	PaidAt          int32              `json:"paidAt"`
	CreatedAt       int32              `json:"createdAt"`
}

var PaymentsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "payments")
