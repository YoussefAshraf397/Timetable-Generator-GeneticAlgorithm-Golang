package Classes

import (
	"go-starter/Models"
)

func ClassTransformer(class Models.Session) map[string]interface{} {
	m := make(map[string]interface{})

	m["class id"] = class.SessionID

	return m

}

func classesTransformer(classes []Models.Session) []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, class := range classes {
		m = append(m, ClassTransformer(class))
	}
	return m
}
