// TODO: implement gorm
// gorm will be better suited for working with json
// and be easier to use
// https://gorm.io/docs/models.html

package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// NOTE: Un-exported struct fields are invisible to the JSON package.
// Export a field by starting it with an UPPERCASE letter.
// Cannot use numbers or symbols
// https://golang.org/ref/spec#Exported_identifiers

type Course struct {
	ID           uuid.UUID   `db:"id" json:"id" validate:"required, uuid"`                   // generated course ID - DB only
	Created_at   time.Time   `db:"created_at" json:"created_at" validate:"required"`         // course created timestamp
	Updated_at   time.Time   `db:"updated_at" json:"updated_at" validate:"required"`         // course updated @ timestamp
	Name         string      `db:"name" json:"name" validate:"required,lte=255"`             // course name e.g., Computer Science Foundations
	CourseNumber string      `db:"course_num" json:"course_num" validate:"required,lte=255"` // course number
	GradeRange   int         `db:"grade_range" json:"grade_range" validate:"required"`       // grades covered
	Active       bool        `db:"active" json:"active" validate:"required"`                 // active status - true|false
	Modules      []string    `db:"modules" json:"modules" validate:"required"`               // modules in course - []string = slice(array) of strings
	CourseAttrs  CourseAttrs `db:"course_attrs" json:"course_attrs" validate:""`
}

type CourseAttrs struct{}

// Value makes LessonAttrs struct implement the driver.Value interface
// method returns a JSON-encoded representation of the struct
func (c CourseAttrs) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan makes LessonAttrs struct implement sql.Scanner interface
// method decodes the JSON-encoded value into struct fields
func (c *CourseAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	// &c points to CourseAttrs
	return json.Unmarshal(j, &c)
}
