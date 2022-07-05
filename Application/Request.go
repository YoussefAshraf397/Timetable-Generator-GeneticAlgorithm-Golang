package Application

import (
	"database/sql"
	"fmt"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-starter/Models"
	"gorm.io/gorm"
)

type Request struct {
	Context         *gin.Context
	DB              *gorm.DB
	Connection      *sql.DB
	IsAuth          bool
	User            *Models.User
	IsAdmin         bool
	Lang            string
	ValidationError error
}

type SharedResources interface {
	Share()
}

// Handle request closure data
func req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDatabase(&request)
		setLang(&request)
		//request.DB, request.Connection = connectToDatabase()

		return &request
	}
}

func (req *Request) ValidateRequest(errors validation.Errors) *Request {
	req.ValidationError = errors.Filter()
	return req
}

func (req *Request) Fails() bool {
	if req.ValidationError != nil {
		req.BadRequest(req.ValidationError)
		return true
	}
	return false
}

func setLang(req *Request) {
	lang := gotrans.DetectLanguage(req.Context.GetHeader("Accept-Language"))
	gotrans.SetDefaultLocale(lang)
	req.Lang = lang
}

func (req *Request) Share() {

}

func (req Request) Response(code int, body map[string]interface{}) {
	CloseConnection(&req)
	req.Context.JSON(code, body)
}

// Initialize new request closure
func NewRequest(c *gin.Context) *Request {
	request := req()
	return request(c)
}

func NewRequestWithAuth(c *gin.Context) *Request {
	return NewRequest(c).Auth()

}

func AuthRequest(c *gin.Context) (*Request, bool) {
	r := NewRequestWithAuth(c)
	if !r.IsAuth {
		r.NoAuth()
		return r, false
	}
	return r, true
}

func (req Request) Auth() *Request {
	req.IsAuth = false
	req.IsAdmin = false
	authHeader := req.Context.GetHeader("Authorized")
	if authHeader != "" {
		fmt.Println("Header: ", authHeader)
		req.DB.Where("token = ?", authHeader).First(&req.User)
		if req.User.ID != 0 {
			req.IsAuth = true
		}
		if req.User.Group == "admin" {
			req.IsAdmin = true
		}
	}
	return &req
}
