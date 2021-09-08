package queries

import (
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
		// return empty obj and error
		return lessons, err
	}
	// hopefully there's query results
	return lessons, nil

}

func (q *LessonQueries) GetLesson(id uuid.UUID) (models.Lesson, error) {
	lesson := models.Lesson{}

	query := `SELECT FROM Lessons WHERE id = $1`

	err := q.Get(&lesson, query, id)
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}

func (q *LessonQueries) CreateLesson() []mode
