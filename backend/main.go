package main

import (
	"os"

	"github.com/edcamero/app-mascota/backend/db"
	myrouter "github.com/edcamero/app-mascota/backend/router"
	"github.com/iris-contrib/middleware/cors"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.UseRouter(crs)
	app.AllowMethods(iris.MethodOptions) // <- permite el Cors

	app.Logger().SetLevel("debug")

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "hacienod ping"})
	})
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/migraciones", func(ctx iris.Context) {
		db.MigrateDB()
	})

	myrouter.AddRutas(app)

	// Listens and serves incoming http requests
	// on http://localhost:8080.
	app.Listen(":" + port)
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
