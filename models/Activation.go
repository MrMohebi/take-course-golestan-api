package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Activation struct {
	Id              primitive.ObjectID   `json:"id" bson:"_id"`
	PlanIds         []primitive.ObjectID `json:"planIds" bson:"planIds"`
	UserId          primitive.ObjectID   `json:"userId" bson:"userId"`
	TotalNumber     int64                `json:"totalNumber" bson:"totalNumber"`
	UsedNumber      int64                `json:"usedNumber" bson:"usedNumber"`
	AdditionalCurse int8                 `json:"additionalCurse" bson:"additionalCurse"`
	CreatedAt       int64                `json:"createdAt" bson:"createdAt"`
}

var ActivationsCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "plans")
