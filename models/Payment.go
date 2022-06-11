package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	UserID          primitive.ObjectID `json:"userID" bson:"userID"`
	OrderId         string             `json:"orderId" bson:"orderId"`
	IdpayId         string             `json:"idpayId" bson:"idpayId"`
	IdpayTrackingId int32              `json:"idpayTrackingId" bson:"idpayTrackingId"`
	PayerName       string             `json:"payerName" bson:"payerName"`
	PayerPhone      string             `json:"payerPhone" bson:"payerPhone"`
	PayerEmail      string             `json:"payerEmail" bson:"payerEmail"`
	CardNumber      string             `json:"cardNumber" bson:"cardNumber"`
	CardNumberHash  string             `json:"cardNumberHash" bson:"cardNumberHash"`
	Link            string             `json:"link" bson:"link"`
	CallbackUrl     string             `json:"callbackUrl" bson:"callbackUrl"`
	Amount          int64              `json:"amount" bson:"amount"`
	PaidAt          int64              `json:"paidAt" bson:"paidAt"`
	VerifiedAt      int64              `json:"verifiedAt" bson:"verifiedAt"`
	CreatedAt       int64              `json:"createdAt" bson:"createdAt"`
}

var PaymentsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "payments")
