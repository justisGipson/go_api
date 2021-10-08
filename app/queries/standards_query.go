package queries

import (
	"fmt"

	"github.com/CodeliciousProduct/bluebird/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StandardsQueries struct {
	*sqlx.DB
}

// TODO: write queries
// this is here so I can finish Queries struct in
// ../../platform/database/open_db_conn.go

func (q *StandardsQueries) GetStandards() ([]models.Standard, error) {
	// define standards
	standards := []models.Standard{}
	// query string
	query := `SELECT * FROM Standards`
	// query db
	err := q.Get(&standards, query)
	if err != nil {
		// return empty standards obj and error message
		return standards, fmt.Errorf("query error: failed to get standards = %v", err)
	}
	return standards, nil
}

func (q *StandardsQueries) GetStandard(id uuid.UUID) (models.Standard, error) {
	standard := models.Standard{}
	// query string
	query := `SELECT FROM Standards WHERE id = $1`
	// query db
	err := q.Get(&standard, query)
	if err != nil {
		return standard, fmt.Errorf("query error: failed to get lesson %c - %v", id, err)
	}
	return standard, nil
}

// Create verbiage felt wrong here...
func (q *StandardsQueries) AddNewStandard(s *models.Standard) (string, error) {
	query := `INSERT INTO standards VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	// send to db
	_, err := q.Exec(
		query,
		s.ID,
		s.Created_at,
		s.Updated_at,
		s.State,
		s.Org,
		s.Grade,
		s.StandardID,
		s.Standard,
		s.Description,
		s.Concept,
		s.Subconcept,
		s.Practice,
	)
	if err != nil {
		// only return error
		return "", fmt.Errorf("query error: failed creating standard - %v", err)
	}
	return fmt.Sprintf("standard %c added", s.ID), nil
}

func (q *StandardsQueries) UpdateStandard(id uuid.UUID, s *models.Standard) (string, error) {
	// query string
	// right now will only update `updated_at` & `standard`
	// I think these fields could be the ones that are updated at any frequency
	query := `UPDATE standards SET updated_at = $2, standard = $7 WHERE id = $1`
	// query and update fields
	_, err := q.Exec(query, id, s.Updated_at, s.Standard)
	if err != nil {
		return "", fmt.Errorf("query error: failed to update standard %c - %v", id, err)
	}
	return fmt.Sprintf("standard %c updated", id), nil
}

func (q *StandardsQueries) DeleteStandard(id uuid.UUID) (string, error) {
	// query string
	query := `DELETE FROM standards WHERE id = $1`
	_, err := q.Exec(query, id)
	if err != nil {
		return "", fmt.Errorf("query error: failed to delete standard %c - %v", id, err)
	}
	return fmt.Sprintf("standard %c deleted", id), nil
}
