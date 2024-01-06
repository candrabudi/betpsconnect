package dto

type (
	ResultAllCoordinatorSubdistrict struct {
		Items    []FindCoordinatorSubdistrict `json:"items"`
		Metadata MetaData                     `json:"metadata"`
	}

	FindCoordinatorSubdistrict struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"kordes_name,omitempty" json:"name"`
		NoHandphone   string `bson:"kordes_phone,omitempty" json:"no_handphone"`
		Nik           string `bson:"kordes_nik,omitempty" json:"nik"`
		Usia          int    `bson:"kordes_age,omitempty" json:"age"`
		Alamat        string `bson:"kordes_address,omitempty" json:"address"`
		NamaKabupaten string `bson:"kordes_city,omitempty" json:"city"`
		NamaKecamatan string `bson:"kordes_district,omitempty" json:"district"`
		Jaringan      string `bson:"kordes_network,omitempty" json:"network"`
	}

	PayloadStoreCoordinatorSubdistrict struct {
		KordesName        string `json:"kordes_name"`
		KordesNik         string `json:"kordes_nik"`
		KordesPhone       string `json:"kordes_phone"`
		KordesAge         int    `json:"kordes_age"`
		KordesAddress     string `json:"kordes_address"`
		KordesCity        string `json:"kordes_city"`
		KordesDistrict    string `json:"kordes_district"`
		KordesSubdistrict string `json:"kordes_subdistrict"`
		KordesNetwork     string `json:"kordes_network"`
	}

	PayloadUpdateCoordinatorSubdistrict struct {
		KordesName        string `json:"kordes_name"`
		KordesNik         string `json:"kordes_nik"`
		KordesPhone       string `json:"kordes_phone"`
		KordesAge         int    `json:"kordes_age"`
		KordesAddress     string `json:"kordes_address"`
		KordesCity        string `json:"kordes_city"`
		KordesDistrict    string `json:"kordes_district"`
		KordesSubdistrict string `json:"kordes_subdistrict"`
		KordesNetwork     string `json:"kordes_network"`
	}

	CoordinationSubdistrictFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Jaringan      string `json:"jaringan"`
		Nama          string `json:"nama"`
	}
)
