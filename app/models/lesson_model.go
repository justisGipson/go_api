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

// TODO: add time/duration field here
type Lesson struct {
	ID                 uuid.UUID `db:"id" json:"id" validate:"required,uuid"`                           // lesson ID
	Created_at         time.Time `db:"created_at" json:"created_at" validate:"required"`                // creation timestamp
	Updated_at         time.Time `db:"updated_at" json:"updated_at" validate:"required"`                // update timestamp
	Name               string    `db:"name" json:"name" validate:"required,lte=255"`                    // lesson name
	LessonNumber       string    `db:"lessonNumber" json:"lessonNumber" validate:"required,lte=255"`    // lesson number
	Course             string    `db:"course" json:"course" validate:"required,lte=255"`                // lesson is part of course
	Active             bool      `db:"active" json:"active" validate:"required"`                        // active & in use: true | false
	CurrentVersion     string    `db:"currentVersion" json:"currentVersion" validate:"required"`        // link to live Google Doc
	GradeRange         int       `db:"gradeRange" json:"gradeRange" validate:"required"`                // grades covered by lesson
	Duration           string    `db:"duration" json:"duration" validate:"required"`                    // lesson length/duration
	LearningObjectives string    `db:"learningObjectives" json:"learningObjectives" validate:"reqired"` // lesson learning objective
	Sel                bool      `db:"sel" json:"sel" validate:"required"`                              // lesson is sel: true | false
	Types              *Types    // lesson type - formerly Pillars
	// standards mapped to lessons k-12
	KStandards      string `db:"k_standards" json:"k_standards"`
	OneStandards    string `db:"1_standards" json:"1_standards"`
	TwoStandards    string `db:"2_standards" json:"2_standards"`
	ThreeStandards  string `db:"3_standards" json:"3_standards"`
	FourStandards   string `db:"4_standards" json:"4_standards"`
	FiveStandards   string `db:"5_standards" json:"5_standards"`
	SixStandards    string `db:"6_standards" json:"6_standards"`
	SevenStandards  string `db:"7_standards" json:"7_standards"`
	EightStandards  string `db:"8_standards" json:"8_standards"`
	NineStandards   string `db:"9_standards" json:"9_standards"`
	TenStandards    string `db:"10_standards" json:"10_standards"`
	ElevenStandards string `db:"11_standards" json:"11_standards"`
	TwelveStandards string `db:"12_standards" json:"12_standards"`
	// dunno bout this one either... will lessons have attributes?
	LessonAttrs LessonAttrs `db:"lesson_attrs" json:"lesson_attrs" validate:""`
}

type Types struct {
	// lesson types - DC, Unplugged, Coding, STEM Career
	Id   uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Name string    `db:"name" json:"name" validate:"required,lte=255"`
	// in Firestore there's a color associated with each type
	// for reference in tables and used as a filter for lesson types
	// Types are formerly "Pillars"
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
