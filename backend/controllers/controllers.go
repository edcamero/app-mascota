package controllers

import (
	"app-mascota/view"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edcamero/app-mascota/backend/db"
	"github.com/edcamero/app-mascota/backend/environment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	response "github.com/edcamero/app-mascota/backend/view"

	"github.com/edcamero/app-mascota/backend/models"
	"github.com/kataras/iris/v12"
)

//GetUser retorna un usuario
func GetUser(ctx iris.Context) {
	//Estructura vacia donde se gurdarán los datos
	users := []models.User{}

	// Se obtiene el parametro id de la URL
	//id := ctx.Params().GetUint64Default("id", 0)

	//Conexión a la DB
	//db := db.GetConnection()

	//Se codifican los datos a formato JSON
	j, _ := json.Marshal(users)
	// Se envian los datos
	response.SendResponse(ctx, http.StatusOK, j)

	// Si no existe se envia un error 404
	log.Println(users)
	//response.SendErr(ctx, http.StatusNotFound)

}

//AllUsers retorna todos los usuarios
func AllUsers(ctx iris.Context) {

	fmt.Println("Conecto con el controlador")
	users := []models.User{}
	db := db.GetConnection()
	filter := bson.D{{}}
	findOptions := options.Find()

	collection := db.Database(environment.DATABASE).Collection("users")
	// Consulta a la DB - SELECT * FROM contacts WHERE ID = ?

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.User

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, elem)
	}
	cur.Close(context.Background())
	//j, _ := json.Marshal(users)
	// Se envian los datos
	ctx.JSON(users)
}

//StoreUsers almacena un usuario
func StoreUsers(ctx iris.Context) {
	// Estructura donde se gurdaran los datos del body
	user := models.User{}
	//db := db.GetConnection()
	//defer db.Close()
	err := json.NewDecoder(ctx.Request().Body).Decode(&user)
	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400
		fmt.Println(err)
		view.SendErr(ctx, http.StatusBadRequest)
		return
	}
	//err = db.Create(&user).Error
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		view.SendErr(ctx, http.StatusInternalServerError)
		return
	}
	// Se codifica el nuevo registro y se devuelve
	j, _ := json.Marshal(user)
	view.SendResponse(ctx, http.StatusCreated, j)

}

//UpdateUser actualizar usuario
func UpdateUser(ctx iris.Context) {
	// Estructuras donde se almacenaran los datos
	//	userFind := models.User{}
	userData := models.User{}
	// Se obtiene el parametro id de la URL
	//id := ctx.Params().Get("id")
	// Conexión a la DB
	//db := db.GetConnection()
	//defer db.Close()
	// Se buscan los datos
	//db.Find(&userFind, id)
	//if userFind.ID > 0 {
	// Si existe el registro se decodifican los datos del body
	err := json.NewDecoder(ctx.Request().Body).Decode(&userData)
	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400
		view.SendErr(ctx, http.StatusBadRequest)
		return
		//	}
		// Se modifican los datos
		//db.Model(&userFind).Updates(userData)
		// Se codifica el registro modificado y se devuelve
		//j, _ := json.Marshal(userFind)
		//view.SendResponse(ctx, http.StatusOK, j)
	} else {
		// Sí no existe el registro especificado se devuelde un error 404
		view.SendErr(ctx, http.StatusNotFound)
	}

}

//DeleteUser  elimina un contacto por ID
func DeleteUser(ctx iris.Context) {
	// Estructura donde se guardara el registo buscado
	//user := models.User{}
	// Se obtiene el parametro id de la URL
	//id := ctx.Params().Get("id")
	//db := db.GetConnection()

	//db.Find(&user, id)

	// Sí existe, se borra y se envia contenido vacio
	//db.Delete(user)
	view.SendResponse(ctx, http.StatusOK, []byte(`{}`))

	// Sí no existe el registro especificado se devuelde un error 404
	view.SendErr(ctx, http.StatusNotFound)

}
