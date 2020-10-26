package router

import (
	Controllers "github.com/edcamero/app-mascota/backend/controllers"
	"github.com/edcamero/app-mascota/backend/environment"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func AddRutas(app *iris.Application) {

	j := jwt.New(jwt.Config{
		// Extract by "token" url parameter.
		//Extractor: jwt.FromParameter("token"),

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(environment.SECRETKEY), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	api := app.Party("/api")
	api.Post("/login", Controllers.Login)
	api.Post("/adoptante/registrar", Controllers.RegisterAdoptante)

	adopApi := api.Party("/adoptante")
	adminApi := api.Party(("/admin"), j.Serve, Controllers.AuthenticatedAdmin)

	adopApi.Get("/", Controllers.AuthenticatedAdoptante)

	adminApi.Post("/registrar", Controllers.RegisterAdmin)

	api.Get("/user/{id:uint64}", Controllers.GetUser)
	api.Get("/user", Controllers.AllUsers)
	api.Post("/user", Controllers.StoreUsers)
	api.Put("/user/{id:uint64}", Controllers.UpdateUser)
	api.Delete("/user/{id:uint64}", Controllers.DeleteUser)

}
