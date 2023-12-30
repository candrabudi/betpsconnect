package model

import "time"

type TrueResident struct {
	ID          int       `bson:"id,omitempty"`
	ResidentID  int       `bson:"resident_id, empty"`
	FullName    string    `bson:"full_name,omitempty"`
	Nik         string    `bson:"nik,omitempty"`
	NoHandphone string    `bson:"no_handphone,omitempty"`
	Age         int       `bson:"age,omitempty"`
	Gender      string    `bson:"gender,omitempty"`
	Address     string    `bson:"address,omitempty"`
	District    string    `bson:"district, omitempty"`
	SubDistrict string    `bson:"subdistrict, omitempty"`
	City        string    `bson:"city, omitempty"`
	BirthDate   string    `bson:"birth_date, omitempty"`
	BirthPlace  string    `bson:"birth_place, omitempty"`
	Tps         string    `bson:"tps, omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}
