package dto

type (
	ResultAllCoordinatorTps struct {
		Items    []FindCoordinatorTps `json:"items"`
		Metadata MetaData             `json:"metadata"`
	}

	FindCoordinatorTps struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"kortps_name,omitempty" json:"name"`
		NoHandphone   string `bson:"kortps_phone,omitempty" json:"no_handphone"`
		Nik           string `bson:"kortps_nik,omitempty" json:"nik"`
		Usia          int    `bson:"kortps_age,omitempty" json:"age"`
		Alamat        string `bson:"kortps_address,omitempty" json:"address"`
		NamaKabupaten string `bson:"kortps_city,omitempty" json:"city"`
		NamaKecamatan string `bson:"kortps_district,omitempty" json:"district"`
		NamaKelurahan string `bson:"kortps_subdistrict,omitempty" json:"subdistrict"`
		Tps           string `bson:"kortps_tps,omitempty" json:"tps"`
		Jaringan      string `bson:"kortps_network,omitempty" json:"network"`
	}

	PayloadStoreCoordinatorTps struct {
		KorTpsName        string `json:"kortps_name"`
		KorTpsNik         string `json:"kortps_nik"`
		KorTpsPhone       string `json:"kortps_phone"`
		KorTpsAge         int    `json:"kortps_age"`
		KorTpsAddress     string `json:"kortps_address"`
		KorTpsCity        string `json:"kortps_city"`
		KorTpsDistrict    string `json:"kortps_district"`
		KorTpsSubdistrict string `json:"kortps_subdistrict"`
		KorTpsTps         string `json:"kortps_tps"`
		KorTpsNetwork     string `json:"kortps_network"`
	}

	PayloadUpdateCoordinatorTps struct {
		KorTpsName        string `json:"kortps_name"`
		KorTpsNik         string `json:"kortps_nik"`
		KorTpsPhone       string `json:"kortps_phone"`
		KorTpsAge         int    `json:"kortps_age"`
		KorTpsAddress     string `json:"kortps_address"`
		KorTpsCity        string `json:"kortps_city"`
		KorTpsDistrict    string `json:"kortps_district"`
		KorTpsSubdistrict string `json:"kortps_subdistrict"`
		KorTpsTps         string `json:"kortps_tps"`
		KorTpsNetwork     string `json:"kortps_network"`
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
