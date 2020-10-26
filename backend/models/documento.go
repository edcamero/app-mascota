package models

type Documento struct {
	Tipo   string `json:"tipo" bson:"tipo"`
	Numero string `json:"numero" bson:"numero"`
}
