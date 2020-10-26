package models

import (
	"time"
)

type Persona struct {
	//ID         primitive.ObjectID `bson:"_id"`
	Nombre     string    `json:"nombre" bson:"nombre"`
	Apellido   string    `json:"apellido" bson:"apellido"`
	Documento  Documento `json:"documento" bson:"documento"`
	Direccion  string    `json:"direccion" bson:"direccion"`
	Telefono   string    `json:"telefono" bson:"telefono"`
	InsertedAt time.Time `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time `json:"last_update" bson:"last_update"`
}
