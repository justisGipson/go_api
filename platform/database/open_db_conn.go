package database

import "github.com/CodeliciousProduct/bluebird/app/queries"

//  queries struct to collect all app queries
type Queries struct {
	*queries.LessonQueries
	*queries.CourseQueries
	*queries.StandardsQueries
	*queries.CurriculumQueries
	*queries.ResourcesQueries
}

func OpenDBConnection() (*Queries, error) {
	// new pgsql connection
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}
	return &Queries{
		// set queries for models
		LessonQueries:     &queries.LessonQueries{DB: db},
		CourseQueries:     &queries.CourseQueries{DB: db},
		StandardsQueries:  &queries.StandardsQueries{DB: db},
		CurriculumQueries: &queries.CurriculumQueries{DB: db},
		ResourcesQueries:  &queries.ResourcesQueries{DB: db},
	}, nil
}
