// TODO: implement an ORM

package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// NOTES: Fields start with an UPPERCASE letter
// otherwise they are invisible to the JSON pkg
// https://golang.org/ref/spec#Exported_identifiers

type Course struct {
	// generated course ID - DB only
	ID uuid.UUID `db:"id" json:"id" validate:"required, uuid"`
	// course created timestamp
	Created_at time.Time `db:"created_at" json:"created_at"`
	// course updated @ timestamp
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
	// course name e.g., Computer Science Foundations
	Name string `db:"name" json:"name" validate:"required,lte=255"`
	// course number
	CourseNumber string `db:"course_num" json:"course_num" validate:"required,lte=255"`
	// grades covered
	GradeRange int `db:"grade_range" json:"grade_range" validate:"required"`
	// active status - true|false
	Active bool `db:"active" json:"active" validate:"required"`
	// modules in course - []byte should work with json
	// which modules will most likely be it's own object
	Modules []byte `db:"modules" json:"modules" validate:"required"`
}

type CourseAttrs struct {
}

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
