package entity

import (
	"strings"
	"time"
)

type Visitor struct {
	Nickname  string    `json:"nickname"`
	VisitTime time.Time `json:"visit_time"`
}

func (v *Visitor) IsValid() bool {
	if len(strings.TrimSpace(v.Nickname)) == 0 {
		return false
	}

	return !v.VisitTime.IsZero()
}
