package Application

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func makeConnection() *gorm.DB {
	dsn := os.Getenv("DATABASE_USER_NAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_USER_HOST") + ":" + os.Getenv("DATABASE_USER_PORT") + ")/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error in database connection !", err.Error())
	}
	return db
}

func returnConnection(db *gorm.DB) *sql.DB {
	connection, err := db.DB()
	if err != nil {
		fmt.Println("Error on database connection", err.Error())
	}

	return connection

}

// Initialize database connection
func connectToDatabase(share SharedResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.DB = makeConnection()
		app.Connection = returnConnection(app.DB)
	case *Request:
		req := share.(*Request)
		req.DB = makeConnection()
		req.Connection = returnConnection(req.DB)
	}

}

// close database connection
func CloseConnection(share SharedResources) {
	switch share.(type) {
	case *Application:
		app := share.(*Application)
		app.Connection.Close()
	case *Request:
		req := share.(*Request)
		req.Connection.Close()
	}
}
