package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserName   string             `json:"username" bson:"username"`
	Password   string             `json:"password" bson:"password"`
	Rol        Rol                `json:"rol" bson:"rol"`
	InsertedAt time.Time          `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time          `json:"last_update" bson:"last_update"`
	//DeleteAt time.Time 		`json:"delete_at" bson:"delete_at"`
	//Edad        uint   `bson:"edad"`
	//Telefono    string `bson:"telefono"`
	//Direccion   string `bson:"direccion"`
	//Email       string `bson:"email"`
	//Descripcion string `bson:"descripcion"`
}
