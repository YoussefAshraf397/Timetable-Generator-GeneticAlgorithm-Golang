package main

import (
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/subosito/gotenv"
	"go-starter/Application"
	"go-starter/Routes"
)

func main() {
	// Gin framework for working with routes  (Routing , validate request , response , middleware)
	// gorm (connect database , quiers)
	// auth user
	// language
	// ------------------------------------------------------------------------------------------------- //

	gotenv.Load(".env")
	app := Application.NewApp()
	_ = gotrans.SetDefaultLocale("en") // Setting default locale
	fmt.Println(gotrans.T("hello_world"))

	// migrate project
	app.Migrate()

	// seeders
	//app.Seed()

	// close connection
	Application.CloseConnection(app)

	routerApp := Routes.RouterApp{app}
	routerApp.Routing()

	// start server app
	app.Gin.Run(":9999") //http://127.0.0.1:9999/ping
}
