package model

import "time"

type CoordinationDistrict struct {
	ID             int       `bson:"id,empty"`
	KorkabID       int       `bson:"korcam_id"`
	KorkabName     string    `bson:"korkab_name,empty"`
	KorcamName     string    `bson:"korcam_name,empty"`
	KorcamNik      string    `bson:"korcam_nik,empty"`
	KorcamPhone    string    `bson:"korcam_phone,empty"`
	KorcamAge      int       `bson:"korcam_age,empty"`
	KorcamAddress  string    `bson:"korcam_address,empty"`
	KorcamDistrict string    `bson:"korcam_district,empty"`
	KorcamCity     string    `bson:"korcam_city,empty"`
	CreatedAt      time.Time `bson:"created_at,empty"`
	UpdatedAt      time.Time `bson:"updated_at,empty"`
}
