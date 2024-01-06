package model

import "time"

type CoordinationSubdistrict struct {
	ID                int       `bson:"id,empty"`
	KordesName        string    `bson:"kordes_name,empty"`
	KordesNik         string    `bson:"kordes_nik,empty"`
	KordesPhone       string    `bson:"kordes_phone,empty"`
	KordesAge         int       `bson:"kordes_age,empty"`
	KordesAddress     string    `bson:"kordes_address,empty"`
	KordesDistrict    string    `bson:"kordes_district,empty"`
	KordesCity        string    `bson:"kordes_city,empty"`
	KordesSubdistrict string    `bson:"kordes_subdistrict,empty"`
	KordesNetwork     string    `bson:"kordes_network,empty"`
	CreatedAt         time.Time `bson:"created_at,empty"`
	UpdatedAt         time.Time `bson:"updated_at,empty"`
}
