package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Phone     string             `json:"phone" bson:"phone"`
	Email     string             `json:"email" bson:"email"`
	Token     string             `json:"token" bson:"token"`
	LastLogin int64              `json:"lastLogin" bson:"lastLogin"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
}

var UsersCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "users")
