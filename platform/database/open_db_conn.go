package database

import "github.com/CodeliciousProduct/bluebird/app/queries"

//  queries struct to collect all app queries
type Queries struct {
	*queries.LessonQueries
}

func OpenDBConnection() (*Queries, error) {
	// new pgsql connection
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}
	return &Queries{
		// set queries from models
		LessonQueries: &queries.LessonQueries{DB: db},
	}, nil
}
