package models

import (
	"errors"
	"strings"
	"time"
)

type LogRequest struct {
	Search string    `json:"search"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

func (l LogRequest) Validate() error {
	if strings.TrimSpace(l.Search) == "" {
		return errors.New("invalid search text")
	}

	return nil
}
