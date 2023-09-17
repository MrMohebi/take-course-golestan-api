package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Plan struct {
	Id                    primitive.ObjectID `json:"id" bson:"_id"`
	Title                 string             `json:"title" bson:"title"`
	Details               string             `json:"details" bson:"details"`
	Price                 int64              `json:"price" bson:"price"`
	IsAdditionalCurseItem bool               `json:"IsAdditionalCurseItem" bson:"IsAdditionalCurseItem"`
	CreatedAt             int64              `json:"createdAt" bson:"createdAt"`
}

var PlansCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "plans")
