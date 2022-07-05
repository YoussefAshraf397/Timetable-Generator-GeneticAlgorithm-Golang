package Routes

import (
	"go-starter/Controllers/Auth"
)

func (app RouterApp) authRoutes() {
	app.Gin.GET("/create-user", Auth.CreateUser)
	app.Gin.GET("/me", Auth.Me)
}
