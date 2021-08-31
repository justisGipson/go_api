package models

import (
	"databases/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	
