package dto

type (
	ResultAllCoordinatorSubdistrict struct {
		Items    []FindCoordinatorSubdistrict `json:"items"`
		Metadata MetaData                     `json:"metadata"`
	}

	FindCoordinatorSubdistrict struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"kordes_name,omitempty" json:"nama"`
		NoHandphone   string `bson:"kordes_phone,omitempty" json:"telp"`
		Nik           string `bson:"kordes_nik,omitempty" json:"nik"`
		Usia          int    `bson:"kordes_age,omitempty" json:"usia"`
		Gender        string `bson:"kordes_gender,omitempty" json:"jenis_kelamin"`
		Alamat        string `bson:"kordes_address,omitempty" json:"alamat"`
		NamaKabupaten string `bson:"kordes_city,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string `bson:"kordes_district,omitempty" json:"nama_kecamatan"`
		NamaKelurahan string `bson:"kordes_subdistrict,omitempty" json:"nama_kelurahan"`
		Jaringan      string `bson:"kordes_network,omitempty" json:"jaringan"`
	}

	PayloadStoreCoordinatorSubdistrict struct {
		KordesName        string `json:"full_name"`
		KordesNik         string `json:"nik"`
		KordesPhone       string `json:"no_handphone"`
		KordesAge         int    `json:"age"`
		KordesGender      string `json:"gender"`
		KordesAddress     string `json:"address"`
		KordesCity        string `json:"city"`
		KordesDistrict    string `json:"district"`
		KordesSubdistrict string `json:"subdistrict"`
		KordesNetwork     string `json:"jaringan"`
	}

	PayloadUpdateCoordinatorSubdistrict struct {
		KordesName        string `json:"full_name"`
		KordesNik         string `json:"nik"`
		KordesPhone       string `json:"no_handphone"`
		KordesAge         int    `json:"age"`
		KordesGender      string `json:"gender"`
		KordesAddress     string `json:"address"`
		KordesCity        string `json:"city"`
		KordesDistrict    string `json:"district"`
		KordesSubdistrict string `json:"subdistrict"`
		KordesNetwork     string `json:"jaringan"`
	}

	CoordinationSubdistrictFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Jaringan      string `json:"jaringan"`
		Nama          string `json:"nama"`
	}
)
