package models

import "time"

type Especie []struct {
	Nombre string `json:"nombre"  bson:"nombre"`
	Razas  []struct {
		Nombre string `json:"nombre"  bson:"nombre"`
	} `json:"razas"  bson:"razas"`
	InsertedAt time.Time `json:"inserted_at" bson:"inserted_at"`
	LastUpdate time.Time `json:"last_update" bson:"last_update"`
}
