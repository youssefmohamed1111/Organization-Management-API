package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
	ID			primitive.ObjectID			`bson:"_id"`
	Name			*string					`json: "name" validate:*required,min=2,max=100`
	email			*string					`json: "password" validate:"required,min=6"`	
	password		*string					`json: "email" validate:'email,required'`
	Token			*string					`json:"token"`
	Refresh_Token 	*string					`json: "refresh_token"`
	Created_at		time.Time				`json: "Created_at"`
	Updated_at		time.Time				`json: "Updated_at"`
	User_id			string					`json: "User_id"`
}
