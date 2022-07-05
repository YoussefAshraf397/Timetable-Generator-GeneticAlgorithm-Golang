package Models

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	Username string `json:"username" grom:"type:varchar(50)"`
	Email    string `json:"email" 	 grom:"type:varchar(50)`
	Password string `json:"password" grom:"type:varchar(50)`
	Token    string `json:"token" 	 grom:"type:varchar(100)`
	Group    string `json:"group" 	 grom:"type:varchar(20)`
}

// another apporch to transform data

func (user User) Transform() map[string]interface{} {
	m := make(map[string]interface{})

	m["email"] = user.Email
	m["username"] = user.Username
	m["token"] = user.Token

	return m
}

func (users Users) Transform() []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, user := range users {
		m = append(m, user.Transform())
	}
	return m
}
