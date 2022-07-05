package Classes

import (
	"go-starter/Models"
)

func ClassTransformer(class Models.Class) map[string]interface{} {
	m := make(map[string]interface{})

	m["class id"] = class.ClassID

	return m

}

func classesTransformer(classes []Models.Class) []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, class := range classes {
		m = append(m, ClassTransformer(class))
	}
	return m
}

