package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID                 uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	created_at         time.Time `db:"created_at" json:"created_at`
	updated_at         time.Time `db:"updated_at" json:"updated_at`
	name               string    `db:"name" json:"name" validate:"required,lte=255"`
	lessonNumber       string    `db:"lessonNumber" json:"lessonNumber" validate:"required,lte=255"`
	course             string    `db:"course" json:"course" validate:"required,lte=255"`
	active             bool      `db:"active" json:"active" validate:"required"`
	currentVersion     string    `db:"currentVersion" json:"currentVersion" validate:"required"`
	gradeRange         int       `db:"gradeRange" json:"gradeRange" validate:"required"`
	learningObjectives string    `db:"learningObjectives" json:"learningObjectives" validate:"required"`
	sel                bool      `db:"sel" json:"sel" validate:"required"`
	types              Types     // ? dunno about this one yet
	kStandards         string    `db:"kStandards" json:"kStandards"`
	oneStandards       string    `db:"oneStandards" json:"oneStandards"`
	twoStandards       string
	threeStandards     string
	fourStandards      string
	fiveStandards      string
	sixStandards       string
	sevenStandards     string
	eightStandards     string
	nineStandards      string
	tenStandards       string
	elevenStandards    string
	twelveStandards    string
}

type Types struct {
	// do these have to be their own thing?
}

// Value makes LessonAttrs struct implement the driver.Value interface
// method returns a JSON-encoded representation of the struct
func (l LessonAttrs) Value() (driver.Value, error) {
	return json.Marshal(l)
}

// Scan makes LessonAttrs struct implement sql.Scanner interface
// method decodes the JSON-encoded value into struct fields
func (l *LessonAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(j, &l)
}
