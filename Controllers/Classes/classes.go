package Classes

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
	"go-starter/Models"
	ClassTransformer "go-starter/Transformers/Classes"
	//"go-starter/Validation/Visitors"
	//"os/user"
)

func CreateSession(c *gin.Context) {
	// Binding Request
	r := Application.NewRequest(c)
	var class Models.Session
	r.Context.ShouldBind(&class)

	//Validate Request
	//if r.ValidateRequest(Visitors.RegisterValidation(user)).Fails() {
	//	return
	//}
	//class.SessionID = 1
	//class.GroupID = 1
	//class.ModuleID = 1
	//class.ProfessorID = 1
	//class.TimeSlotID = 1
	//class.RoomID = 1
	r.DB.Create(&class)
	r.Created(ClassTransformer.ClassTransformer(class))
}
