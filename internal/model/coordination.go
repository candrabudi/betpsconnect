package model

import "time"

type CoordinationCity struct {
	ID                      int       `bson:"id,empty"`
	CoordinationName        string    `bson:"nama_kabupaten,empty"`
	CoordinationPhone       string    `bson:"nama_kecamatan,empty"`
	CoordinationAge         string    `bson:"nama_kecamatan,empty"`
	CoordinationAddress     string    `bson:"nama_kecamatan,empty"`
	CoordinationSubdistrict string    `bson:"nama_kecamatan,empty"`
	CoordinationDistrict    string    `bson:"nama_kecamatan,empty"`
	CreatedAt               time.Time `bson:"nama_kecamatan,empty"`
	UpdatedAt               time.Time `bson:"nama_kecamatan,empty"`
}
