package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID primitive.ObjectID `bson:"_id"`
	//Persona Persona            `json:"nombre" bson:"rolnombre"`
	User   User               `json:"user" bson:"-"`
	UserId primitive.ObjectID `json:"-" bson:"userId"`
	Persona
}
