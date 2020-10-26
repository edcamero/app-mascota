package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/edcamero/app-mascota/backend/db"
	"github.com/edcamero/app-mascota/backend/environment"
	"github.com/edcamero/app-mascota/backend/models"
	"github.com/edcamero/app-mascota/backend/util"
	"github.com/edcamero/app-mascota/backend/view"
)

func RegisterAdoptante(ctx iris.Context) {
	user := models.User{}
	rol := models.Rol{Nombre: "adoptante"}
	err := json.NewDecoder(ctx.Request().Body).Decode(&user)
	user.ID = primitive.NewObjectID()
	user.InsertedAt = time.Now()
	user.LastUpdate = time.Now()
	fmt.Println(time.Now())
	user.Rol = rol

	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400

		log.Fatal(err)
		//view.SendErr(ctx, http.StatusBadRequest)
		return
	}
	password := util.Encrypt([]byte(user.Password))
	user.Password = password
	//fmt.Println(user)
	conexion := db.GetConnection()
	collection := conexion.Database(environment.DATABASE).Collection("users")
	_, err = collection.InsertOne(context.TODO(), &user)
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		//fmt.Println(err)
		log.Println(err)
		//ctx.Problem(err)
		//ctx.JSON(err)
		ctx.Problem(view.NewProductProblem("Error en la Base de Datos", "correo ya existe"), iris.ProblemOptions{

			JSON: iris.JSON{
				Indent: "  ",
			},

			RetryAfter: 300,
		})
	} else {
		ctx.JSON(user)
	}
	// Se codifica el nuevo registro y se devuelve

}

func RegisterAdmin(ctx iris.Context) {
	admin := models.Admin{}
	rol := models.Rol{Nombre: "admin"}
	err := json.NewDecoder(ctx.Request().Body).Decode(&admin)
	admin.ID = primitive.NewObjectID()
	admin.InsertedAt = time.Now()
	admin.LastUpdate = time.Now()
	fmt.Println(time.Now())
	admin.User.Rol = rol

	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400

		log.Fatal(err)
		//view.SendErr(ctx, http.StatusBadRequest)
		return
	}
	password := util.Encrypt([]byte(admin.User.Password))
	admin.User.Password = password
	//fmt.Println(user)
	conexion := db.GetConnection()
	collection := conexion.Database(environment.DATABASE).Collection("users")
	admin.User.ID = primitive.NewObjectID()
	admin.UserId = admin.User.ID
	_, err = collection.InsertOne(context.TODO(), &admin.User)
	collection = conexion.Database(environment.DATABASE).Collection("administradores")
	_, err = collection.InsertOne(context.TODO(), &admin)
	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		//fmt.Println(err)
		log.Println(err)
		//ctx.Problem(err)
		//ctx.JSON(err)
		ctx.Problem(view.NewProductProblem("Error en la Base de Datos", "correo ya existe"), iris.ProblemOptions{

			JSON: iris.JSON{
				Indent: "  ",
			},

			RetryAfter: 300,
		})
	} else {

		ctx.JSON(admin)
	}
	// Se codifica el nuevo registro y se devuelve

}
