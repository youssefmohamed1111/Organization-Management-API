package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
	ID			primitive.ObjectID			`bson:"_id"`
	Name			*string					`json: "name" validate:*required,min=2,max=100`
	Email			*string					`json: "password" validate:"required,min=6"`	
	Password		*string					`json: "email" validate:'email,required'`
	Token			*string					`json:"token"`
	Refresh_Token 	*string					`json: "refresh_token"`
	User_Type 		*string					`json: "user-type"`
	Created_at		time.Time				`json: "created_at"`
	Updated_at		time.Time				`json: "updated_at"`
	User_id			string					`json: "User_id"`
}
