package models

import "github.com/fatih/structs"

type Class struct {
	ID               string `json:"id" structs:"id" bson:"_id" db:"id"`
	Name             string `json:"name" structs:"name" bson:"name" db:"name"`
	StudentsEnrolled int    `json:"studentsEnrolled" structs:"studentsEnrolled" bson:"studentsEnrolled" db:"studentsEnrolled"`
}

// Map converts structs to a map representation
func (c *Class) Map() map[string]interface{} {
	return structs.Map(c)
}

func (c *Class) Names() []string {
	fields := structs.Fields(c)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
