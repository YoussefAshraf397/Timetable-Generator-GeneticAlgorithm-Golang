package Subjects

import "go-starter/Models"

func SubjectTransformer(user Models.Subject) map[string]interface{} {
	m := make(map[string]interface{})

	//m["email"] = user.Email
	//m["username"] = user.Username
	//m["token"] = user.Token

	return m

}
