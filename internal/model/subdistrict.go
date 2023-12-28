package model

type SubDistrict struct {
	ID            int    `bson:"id,empty"`
	NamaKabupaten string `bson:"nama_kabupaten,empty"`
	NamaKecamatan string `bson:"nama_kecamatan,empty"`
	NamaKelurahan string `bson:"nama_kelurahan,empty"`
}
