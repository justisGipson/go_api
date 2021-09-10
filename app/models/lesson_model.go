package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Un-exported struct fields are invisible to the JSON package.
// Export a field by starting it with an uppercase letter.
// https://golang.org/ref/spec#Exported_identifiers

type Lesson struct {
	// lesson ID
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	// creation timestamp
	Created_at time.Time `db:"created_at" json:"created_at"`
	// update timestamp
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
	// lesson name
	Name string `db:"name" json:"name" validate:"required,lte=255"`
	// lesson number
	LessonNumber string `db:"lessonNumber" json:"lessonNumber" validate:"required,lte=255"`
	// lesson is part of course
	Course string `db:"course" json:"course" validate:"required,lte=255"`
	// active & in use: true | false
	Active bool `db:"active" json:"active" validate:"required"`
	// link to live Google Doc
	CurrentVersion string `db:"currentVersion" json:"currentVersion" validate:"required"`
	// grades covered by lesson
	GradeRange int `db:"gradeRange" json:"gradeRange" validate:"required"`
	// lesson learning objective
	LearningObjectives string `db:"learningObjectives" json:"learningObjectives" validate:"reqired"`
	// lesson is sel: true | false
	Sel bool `db:"sel" json:"sel" validate:"required"`
	// dunno yet
	// Types Types
	// standards mapped to lessons k-12
	KStandards      string `db:"kStandards" json:"kStandards"`
	OneStandards    string `db:"oneStandards" json:"oneStandards"`
	TwoStandards    string
	ThreeStandards  string
	FourStandards   string
	FiveStandards   string
	SixStandards    string
	SevenStandards  string
	EightStandards  string
	NineStandards   string
	TenStandards    string
	ElevenStandards string
	TwelveStandards string
	// dunno bout this one either... will lessons have attributes?
	LessonAttrs LessonAttrs `db:"lesson_attrs" json:"lesson_attrs" validate:""`
}

type Types struct {
	// do these have to be their own thing?
}

// Lesson Attributes...
type LessonAttrs struct {
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
	// &l points to LessonAttrs
	return json.Unmarshal(j, &l)
}
