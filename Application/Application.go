package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"go-starter/Database"
	"go-starter/Seeders"
	"gorm.io/gorm"
)

type Application struct {
	Gin        *gin.Engine
	DB         *gorm.DB
	Connection *sql.DB
}

func (app *Application) Share() {

}

func App() func() *Application {
	return func() *Application {
		var application Application
		application.Gin = gin.Default()
		connectToDatabase(&application)
		err := gotrans.InitLocales("public/languages") // initialize the language files
		if err != nil {
			fmt.Println("Error in loading translation files !", err.Error())
		}

		return &application

	}
}

// Initialize new request closure
func NewApp() *Application {
	app := App()
	return app()
}

func (app *Application) Seed() {
	Seeders.Seeds(app.DB)
}

func (app *Application) Migrate() {
	Database.Migrate(app.DB)
}
