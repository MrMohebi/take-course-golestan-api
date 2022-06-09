package models

import (
	"github.com/MrMohebi/take-course-golestan-api.git/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Phone     string             `json:"phone"`
	Email     string             `json:"email"`
	Token     string             `json:"token"`
	LastLogin int64              `json:"lastLogin"`
	CreatedAt int64              `json:"createdAt"`
}

var UsersCollection *mongo.Collection = configs.GetCollection(configs.GetDBClint(), "users")
