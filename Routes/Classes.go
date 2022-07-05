package Routes

import (
	"go-starter/Controllers/Classes"
)

func (app RouterApp) classesRoutes() {
	app.Gin.POST("/classes", Classes.CreateSession)
}
