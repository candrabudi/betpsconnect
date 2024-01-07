package dto

import "time"

type (
	ResultAllCoordinatorTps struct {
		Items    []FindCoordinatorTps `json:"items"`
		Metadata MetaData             `json:"metadata"`
	}

	FindCoordinatorTps struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"kortps_name,omitempty" json:"nama"`
		NoHandphone   string `bson:"kortps_phone,omitempty" json:"telp"`
		Nik           string `bson:"kortps_nik,omitempty" json:"nik"`
		Usia          int    `bson:"kortps_age,omitempty" json:"usia"`
		Gender        string `bson:"kortps_gender,omitempty" json:"jenis_kelamin"`
		Alamat        string `bson:"kortps_address,omitempty" json:"alamat"`
		NamaKabupaten string `bson:"kortps_city,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string `bson:"kortps_district,omitempty" json:"nama_kecamatan"`
		NamaKelurahan string `bson:"kortps_subdistrict,omitempty" json:"nama_kelurahan"`
		Tps           string `bson:"kortps_tps,omitempty" json:"tps"`
		Jaringan      string `bson:"kortps_network,omitempty" json:"jaringan"`
	}

	ExportCoordinatorTps struct {
		ID            int       `bson:"id,omitempty" json:"id"`
		Nama          string    `bson:"kortps_name,omitempty" json:"nama"`
		NoHandphone   string    `bson:"kortps_phone,omitempty" json:"telp"`
		Nik           string    `bson:"kortps_nik,omitempty" json:"nik"`
		Usia          int       `bson:"kortps_age,omitempty" json:"usia"`
		Gender        string    `bson:"kortps_gender,omitempty" json:"jenis_kelamin"`
		Alamat        string    `bson:"kortps_address,omitempty" json:"alamat"`
		NamaKabupaten string    `bson:"kortps_city,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string    `bson:"kortps_district,omitempty" json:"nama_kecamatan"`
		NamaKelurahan string    `bson:"kortps_subdistrict,omitempty" json:"nama_kelurahan"`
		Tps           string    `bson:"kortps_tps,omitempty" json:"tps"`
		Jaringan      string    `bson:"kortps_network,omitempty" json:"jaringan"`
		CreatedAt     time.Time `bson:"created_at,omitempty" json:"created_at"`
		UpdatedAt     time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	}

	PayloadStoreCoordinatorTps struct {
		KorTpsName        string `json:"full_name"`
		KorTpsNik         string `json:"nik"`
		KorTpsPhone       string `json:"no_handphone"`
		KorTpsAge         int    `json:"age"`
		KorTpsGender      string `json:"gender"`
		KorTpsAddress     string `json:"address"`
		KorTpsCity        string `json:"city"`
		KorTpsDistrict    string `json:"district"`
		KorTpsSubdistrict string `json:"subdistrict"`
		KorTpsTps         string `json:"tps"`
		KorTpsNetwork     string `json:"jaringan"`
	}

	PayloadUpdateCoordinatorTps struct {
		KorTpsName        string `json:"full_name"`
		KorTpsNik         string `json:"nik"`
		KorTpsPhone       string `json:"no_handphone"`
		KorTpsAge         int    `json:"age"`
		KorTpsGender      string `json:"gender"`
		KorTpsAddress     string `json:"address"`
		KorTpsCity        string `json:"city"`
		KorTpsDistrict    string `json:"district"`
		KorTpsSubdistrict string `json:"subdistrict"`
		KorTpsTps         string `json:"tps"`
		KorTpsNetwork     string `json:"jaringan"`
	}

	CoordinationTpsFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Jaringan      string `json:"jaringan"`
		Tps           string `json:"tps"`
		Nama          string `json:"nama"`
	}
)
