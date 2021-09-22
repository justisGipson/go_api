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
	Created_at time.Time
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
