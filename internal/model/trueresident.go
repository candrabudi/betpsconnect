package model

import "time"

type TrueResident struct {
	ID          int       `bson:"id,omitempty"`
	FullName    string    `bson:"full_name,omitempty"`
	Nik         string    `bson:"nik,omitempty"`
	NoHandphone string    `bson:"no_handphone,empty"`
	Age         int       `bson:"age,omitempty"`
	Gender      string    `bson:"gender,omitempty"`
	Address     string    `bson:"address,empty"`
	SubDistrict string    `bson:"subdistrict, omitempty"`
	District    string    `bson:"district, omitempty"`
	City        string    `bson:"city, omitempty"`
	Tps         string    `bson:"tps, omitempty"`
	KorcamName  string    `bson:"korcam_name, empty"`
	KordesName  string    `bson:"kordes_name, empty"`
	KortpsName  string    `bson:"kortps_name, empty"`
	Jaringan    string    `bson:"network, empty"`
	IsManual    int       `bson:"is_manual, empty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty"`
}
