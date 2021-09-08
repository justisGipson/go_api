package queries

import (
	"fmt"

	"github.com/CodeliciousProduct/bluebird/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LessonQueries struct {
	*sqlx.DB
}

// Get lessons method - get all
func (q *LessonQueries) GetLessons() ([]models.Lesson, error) {
	// define lessons var
	lessons := []models.Lesson{}
	// define query string
	query := `SELECT * FROM Lessons`
	// query db
	err := q.Get(&lessons, query)
	if err != nil {
		// return empty obj and error message
		return lessons, fmt.Errorf("query error: failed to get lessons - %e", err)
	}
	// hopefully there's query results
	return lessons, nil

}

func (q *LessonQueries) GetLesson(id uuid.UUID) (models.Lesson, error) {
	lesson := models.Lesson{}

	query := `SELECT FROM Lessons WHERE id = $1`

	err := q.Get(&lesson, query, id)
	if err != nil {
		// return empty object and error message
		return lesson, fmt.Errorf("query error: failed to get lesson - %e", err)
	}
	return lesson, nil
}

func (q *LessonQueries) CreateLesson(l *models.Lesson) error {
	// query string for creating book
	query := `INSERT INTO lessons VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)`
	// send to DB, cross fingers
	_, err := q.Exec(query, l.Created_at, l.Updated_at, l.Name, l.LessonNumber, l.Course, l.Active, l.CurrentVersion, l.GradeRange, l.LearningObjectives, l.Sel, l.KStandards, l.OneStandards, l.TwoStandards, l.ThreeStandards, l.FourStandards, l.FiveStandards, l.SixStandards, l.SevenStandards, l.EightStandards, l.NineStandards, l.TenStandards, l.ElevenStandards, l.TwelveStandards, l.LessonAttrs)
	if err != nil {
		// only returning error
		return fmt.Errorf("query error: failed creating course - %e", err)
	}
	// query isn't meant to return anything
	return nil
}
