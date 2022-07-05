package Visitors

import (
	"go-starter/Models"
)

func UsertTransformer(user Models.User) map[string]interface{} {
	m := make(map[string]interface{})

	m["email"] = user.Email
	m["username"] = user.Username
	m["token"] = user.Token

	return m

}

func UsersTransformer(users []Models.User) []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, user := range users {
		m = append(m, UsertTransformer(user))
	}
	return m
}
