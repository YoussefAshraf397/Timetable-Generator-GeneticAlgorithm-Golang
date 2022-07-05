package Visitors

import (
	"github.com/gin-gonic/gin"
	"go-starter/Application"
	"go-starter/Models"
	VisitorTrasformer "go-starter/Transformers/Visitors"
	"go-starter/Validation/Visitors"
)

func Register(c *gin.Context) {
	// Binding Request
	r := Application.NewRequest(c)
	var user Models.User
	r.Context.ShouldBind(&user)

	//Validate Request
	if r.ValidateRequest(Visitors.RegisterValidation(user)).Fails() {
		return
	}
	user.Token = user.Email
	user.Group = "user"
	r.DB.Create(&user)
	r.Created(VisitorTrasformer.UsertTransformer(user))
}

func Login(c *gin.Context) {
	r := Application.NewRequest(c)
	var user Models.User
	r.Context.ShouldBind(&user)
	if r.ValidateRequest(Visitors.LoginValidation(user)).Fails() {
		return
	}

	r.DB.Where("email =?", user.Email).Where("password =?", user.Password).First(&user)
	if user.ID == 0 {
		r.UserNotFound()
		return
	}
	//r.OK(VisitorTrasformer.UsertTransformer(user)) 	 //#1
	r.OK(user.Transform()) //2
}
