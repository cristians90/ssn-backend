package models

import "time"

type BaseModel struct {
	ID         uint64 `json:"id" storm:"id,increment"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DisabledAt time.Time
	Enabled    bool
}
