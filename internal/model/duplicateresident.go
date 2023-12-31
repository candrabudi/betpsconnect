package model

import "time"

type (
	DuplucateResident struct {
		ID         int       `bson:"id,empty"`
		ResidentID int       `bson:"resident_id,empty"`
		Nik        int       `nik:"id,empty"`
		CreatedAt  time.Time `bson:"created_at,empty"`
		UpdateAt   time.Time `bson:"updated_at,empty"`
	}
)
