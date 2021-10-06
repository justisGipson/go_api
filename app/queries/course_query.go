// db queries for lessons
// Pure SQL queries for now,
// will implement`gorm` in the near future

package queries

import (
	"fmt"

	"github.com/CodeliciousProduct/bluebird/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CourseQueries struct {
	*sqlx.DB
}

func (q *CourseQueries) GetCourses() ([]models.Course, error) {
	// define course var
	courses := []models.Course{}
	// query string
	query := `SELECT * FROM Courses`
	// db query
	// &courses pointer
	err := q.Get(&courses, query)
	if err != nil {
		// return empty course obj and error
		return courses, fmt.Errorf("query error: failed to get courses: %v", err)
	}
	// return all courses
	return courses, nil
}

func (q *CourseQueries) GetCourse(id uuid.UUID) (models.Course, error) {
	course := models.Course{}
	query := `SELECT FROM Courses WHERE id = $1`
	err := q.Get(&course, query, id)
	if err != nil {
		return course, fmt.Errorf("query error: failed to get course with id: %v", err)
	}
	return course, nil
}

func (q *CourseQueries) CreateNewCourse(c *models.Course) (string, error) {
	query := `INSERT INTO courses VALUES($1, $2, $3, $4. $5. $6, $7, $8, $9)`
	_, err := q.Exec(
		query,
		c.ID,
		c.Created_at,
		c.Updated_at,
		c.Name,
		c.CourseNumber,
		c.GradeRange,
		c.Active,
		c.Modules,
		c.CourseAttrs,
	)
	if err != nil {
		return "", fmt.Errorf("query error: failed to create course - %v", err)
	}
	return fmt.Sprintf("course %c created", c.ID), nil
}

func (q *CourseQueries) UpdateCourse(id uuid.UUID, c *models.Course) (string, error) {
	// query string - updates course `updated_at`, `modules`
	// may have to rethink what gets updated, or what could get updated
	// going to get some time w/ PM to discuss possibilities
	query := `UPDATE lessons SET updated_at = $3, modules = $8 WHERE id = $1`
	_, err := q.Exec(query, id, c.Updated_at, c.Modules)
	if err != nil {
		// return empty string and error if there is one
		return "", fmt.Errorf("query error: failed to update course = %v", err)
	}
	return fmt.Sprintf("course %c updated", id), nil
}

func (q *CourseQueries) DeleteCourse(id uuid.UUID) (string, error) {
	query := `DELETE FROM courses WHERE id = $1`
	_, err := q.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("query error: failed to delete course - %v", err)
	}
	return fmt.Sprintf("course %c deleted", id), nil

}
