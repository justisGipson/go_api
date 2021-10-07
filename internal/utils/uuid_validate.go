package utils

import (
	_ "github.com/CodeliciousProduct/bluebird/app/models"
	"github.com/google/uuid"
)

// uuid.Parse is was faster than using regex
// like ~18x faster...
func validateUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
