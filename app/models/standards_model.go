// TODO: implement gorm
// gorm will be better suited for working with json
// and be easier to use
// https://gorm.io/docs/models.html

package models

import (
	"time"

	"github.com/google/uuid"
)

// NOTE: Un-exported struct fields are invisible to the JSON package.
// Export a field by starting it with an UPPERCASE letter.
// Cannot use numbers or symbols
// https://golang.org/ref/spec#Exported_identifiers

// For all standards - state and national
type Standard struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Created_at  time.Time `db:"created_at" json:"created_at" validate:"required"`
	Updated_at  time.Time `db:"updated_at" json:"updated_at" validate:"required"`
	State       string    `db:"state" json:"state" validate:"lte=255"`
	Org         string    `db:"org" json:"org" validate:"lte=255"`
	Grade       string    `db:"grade" json:"grade" validate:"required"`
	StandardID  string    `db:"standard_id" json:"standard_id" validate:"required"`
	Standard    string    `db:"standard" json:"standard" validate:"required"`
	Description string    `db:"description" json:"description"`
	Concept     string    `db:"concept" json:"concept"`
	Subconcept  string    `db:"subconcept" json:"subconcept"`
	Practice    string    `db:"practice" json:"practice"`
}

// don't think we need to implement a StandardsAttr Struct
// but if we do, we need to implement Value and Scan funcs
// to implement driver.Value & sql.Scanner for encoding/decoding
// of json representation of the structs
// check lesson_model for example
