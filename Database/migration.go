package Database

import (
	"go-starter/Models"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&Models.User{},
		&Models.Session{},
		&Models.Subject{},
		&Models.Teacher{},
		&Models.TimeSlot{},
		&Models.Classroom{},
	)

}
