package Routes

import "go-starter/Application"

type RouterApp struct {
	*Application.Application
}

func (app RouterApp) Routing() {
	app.visitorsRoutes()
	app.authRoutes()
	app.adminsRoutes()
	app.classesRoutes()
}
