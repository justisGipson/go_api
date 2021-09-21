// db queries for lessons
// Pure SQL queries for now, will implement something like `gorm`
// when I optimize

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

// *lessonQueries being set through pointer to queries for all queries
func (q *LessonQueries) GetLessons() ([]models.Lesson, error) {
	// define lessons var
	lessons := []models.Lesson{}
	// define query string
	query := `SELECT * FROM Lessons`
	// query db
	// &lessons pointer
	err := q.Get(&lessons, query)
	if err != nil {
		// return empty obj and error message
		return lessons, fmt.Errorf("query error: failed to get lessons - %s", err)
	}
	// hopefully there's query results, and it's all the lessons
	return lessons, nil

}

func (q *LessonQueries) GetLesson(id uuid.UUID) (models.Lesson, error) {
	lesson := models.Lesson{}
	// query string
	query := `SELECT FROM Lessons WHERE id = $1`
	// query db
	err := q.Get(&lesson, query, id)
	if err != nil {
		// return empty object and error message
		return lesson, fmt.Errorf("query error: failed to get lesson - %s", err)
	}
	// return single lesson
	return lesson, nil
}

func (q *LessonQueries) CreateLesson(l *models.Lesson) (string, error) {
	// query string for creating lesson
	query := `INSERT INTO lessons VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)`
	// send to DB, cross fingers
	_, err := q.Exec(query, l.Created_at, l.Updated_at, l.Name, l.LessonNumber, l.Course, l.Active, l.CurrentVersion, l.GradeRange, l.LearningObjectives, l.Sel, l.KStandards, l.OneStandards, l.TwoStandards, l.ThreeStandards, l.FourStandards, l.FiveStandards, l.SixStandards, l.SevenStandards, l.EightStandards, l.NineStandards, l.TenStandards, l.ElevenStandards, l.TwelveStandards, l.LessonAttrs)
	if err != nil {
		// only returning error
		return "", fmt.Errorf("query error: failed creating course - %s", err)
	}
	// query isn't meant to return anything
	// TODO: implement string formatting for this
	return "course created", nil
}

func (q *LessonQueries) UpdateLesson(id uuid.UUID, l *models.Lesson) (string, error) {
	// query string
	// right now updates `updated_at`, `active`
	// need to figure out all fields that could be updated, probably all...
	query := `UPDATE lessons SET updated_at = $2, active = $6 WHERE id = $1`
	// query db, update fields
	_, err := q.Exec(query, id, l.Updated_at, l.Active)
	if err != nil {
		// return err message
		return "", fmt.Errorf("query error: failed to update course - %s", err)
	}
	// return nothing
	return fmt.Sprintf("course %c updated", id), nil
}

func (q *LessonQueries) DeleteLesson(id uuid.UUID) (string, error) {
	// query string
	query := `DELETE FROM lessons WHERE id = $1`
	// send to db
	_, err := q.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("query error: failed to delete lesson - %s", err)
	}
	return fmt.Sprintf("course %c deleted", id), nil
}
