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

type Curriculum struct {
	// generated Curriculum ID
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	// curriculum created timestamp
	Created_at time.Time `db:"created_at" json:"created_at" validate:"required"`
	// curriculum updated timestamp
	Updated_at time.Time `db:"updated_at" json:"updated_at" validate:"required"`
	// date curriculum license starts
	StartDate time.Time `db:"start_date" json:"start_date" validate:"required"`
	// date curriculum license ends
	EndDate time.Time `db:"end_date" json:"end_date" validate:"required"`
	// name of customer delivery
	Name string `db:"name" json:"name" validate:"required,lte=255"`
	// course title e.g, Computer Science Foundations
	CourseTitle string `db:"course_title" json:"course_title" validate:"required,lte=255"`
	// grades covered
	Grades       []byte `db:"grades" json:"grades" validate:"required"`
	Documents    Documents
	Modules      []byte `db:"modules" json:"modules" validate:"required"`
	ModuleLength int    `db:"module_length" json:"module_length" validate:"required"`
	NumModules   int    `db:"num_modules" json:"num_modules" validate:"required"`
	Types               // figure this out
	// indicate selected grades covered in curriculum, this may not need to be a thing
	// especially since course are grade-banded
	CheckedK  bool `db:"checkedK" json:"checkedK"`
	Checked1  bool `db:"checked1" json:"checked1"`
	Checked2  bool `db:"checked2" json:"checked2"`
	Checked3  bool `db:"checked3" json:"checked3"`
	Checked4  bool `db:"checked4" json:"checked4"`
	Checked5  bool `db:"checked5" json:"checked5"`
	Checked6  bool `db:"checked6" json:"checked6"`
	Checked7  bool `db:"checked7" json:"checked7"`
	Checked8  bool `db:"checked8" json:"checked8"`
	Checked9  bool `db:"checked9" json:"checked9"`
	Checked10 bool `db:"checked10" json:"checked10"`
	Checked11 bool `db:"checked11" json:"checked11"`
	Checked12 bool `db:"checked12" json:"checked12"`
}

type Documents struct {
	// need to tie this to the lesson_model somehow
	// since documents are lessons built into a curriculum
}

type CurriculumAttrs struct {
}

func (c CurriculumAttrs) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CurriculumAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(j, &c)
}
