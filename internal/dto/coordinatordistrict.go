package dto

type (
	ResultAllCoordinatorDistrict struct {
		Items    []FindCoordinatorDistrict `json:"items"`
		Metadata MetaData                  `json:"metadata"`
	}

	FindCoordinatorDistrict struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"korcam_name,omitempty" json:"name"`
		NoHandphone   string `bson:"korcam_phone,omitempty" json:"no_handphone"`
		Nik           string `bson:"korcam_nik,omitempty" json:"nik"`
		Usia          int    `bson:"korcam_age,omitempty" json:"age"`
		Alamat        string `bson:"korcam_address,omitempty" json:"address"`
		NamaKabupaten string `bson:"korcam_city,omitempty" json:"city"`
		NamaKecamatan string `bson:"korcam_district,omitempty" json:"district"`
		Jaringan      string `bson:"korcam_network,omitempty" json:"network"`
	}

	PayloadStoreCoordinatorDistrict struct {
		KorcamName     string `json:"korcam_name"`
		KorcamNik      string `json:"korcam_nik"`
		KorcamPhone    string `json:"korcam_phone"`
		KorcamAge      int    `json:"korcam_age"`
		KorcamAddress  string `json:"korcam_address"`
		KorcamCity     string `json:"korcam_city"`
		KorcamDistrict string `json:"korcam_district"`
		KorcamNetwork  string `json:"korcam_network"`
	}

	PayloadUpdateCoordinatorDistrict struct {
		KorcamName     string `json:"korcam_name"`
		KorcamNik      string `json:"korcam_nik"`
		KorcamPhone    string `json:"korcam_phone"`
		KorcamAge      int    `json:"korcam_age"`
		KorcamAddress  string `json:"korcam_address"`
		KorcamCity     string `json:"korcam_city"`
		KorcamDistrict string `json:"korcam_district"`
		KorcamNetwork  string `json:"korcam_network"`
	}

	CoordinationDistrictFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Jaringan      string `json:"jaringan"`
		Nama          string `json:"nama"`
	}
)
