package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActiveCode struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"userID" bson:"userID"`
	Code       string             `json:"code" bson:"code"`
	Phone      string             `json:"phone" bson:"phone"`
	DeviceHash string             `json:"deviceHash" bson:"deviceHash"`
	CreatedAt  int64              `json:"createdAt" bson:"createdAt"`
}

var ActiveCodesCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "activeCodes")
