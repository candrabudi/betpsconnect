package model

import "time"

type CoordinationTps struct {
	ID                int       `bson:"id,empty"`
	KorTpsName        string    `bson:"kortps_name,empty"`
	KorTpsNik         string    `bson:"kortps_nik,empty"`
	KorTpsPhone       string    `bson:"kortps_phone,empty"`
	KorTpsAge         int       `bson:"kortps_age,empty"`
	KorTpsAddress     string    `bson:"kortps_address,empty"`
	KorTpsDistrict    string    `bson:"kortps_district,empty"`
	KorTpsCity        string    `bson:"kortps_city,empty"`
	KorTpsSubdistrict string    `bson:"kortps_subdistrict,empty"`
	KorTpsTps         string    `bson:"kortps_tps,empty"`
	KorTpsNetwork     string    `bson:"kortps_network,empty"`
	CreatedAt         time.Time `bson:"created_at,empty"`
	UpdatedAt         time.Time `bson:"updated_at,empty"`
}
