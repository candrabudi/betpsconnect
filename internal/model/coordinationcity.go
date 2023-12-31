package model

import "time"

type CoordinationCity struct {
	ID            int       `bson:"id,empty"`
	KorkabName    string    `bson:"korkab_name,empty"`
	KorkabNik     string    `bson:"korkab_nik,empty"`
	KorkabPhone   string    `bson:"korkab_phone,empty"`
	KorkabAge     int       `bson:"korkab_age,empty"`
	KorkabGender  string    `bson:"korkab_gender,empty"`
	KorkabAddress string    `bson:"korkab_address,empty"`
	KorkabCity    string    `bson:"korkab_city,empty"`
	KorkabNetwork string    `bson:"korkab_network,empty"`
	CreatedAt     time.Time `bson:"created_at,empty"`
	UpdatedAt     time.Time `bson:"updated_at,empty"`
}
