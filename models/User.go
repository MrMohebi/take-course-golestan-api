package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Phone     string             `json:"phone" bson:"phone"`
	Email     string             `json:"email" bson:"email"`
	Token     string             `json:"token" bson:"token"`
	Code      string             `json:"code" bson:"code"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
}

var UsersCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "users")
