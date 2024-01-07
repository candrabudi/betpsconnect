package dto

type (
	ResultAllCoordinatorDistrict struct {
		Items    []FindCoordinatorDistrict `json:"items"`
		Metadata MetaData                  `json:"metadata"`
	}

	FindCoordinatorDistrict struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"korcam_name,omitempty" json:"nama"`
		NoHandphone   string `bson:"korcam_phone,omitempty" json:"telp"`
		Nik           string `bson:"korcam_nik,omitempty" json:"nik"`
		Usia          int    `bson:"korcam_age,omitempty" json:"usia"`
		Gender        string `bson:"korcam_gender,omitempty" json:"jenis_kelamin"`
		Alamat        string `bson:"korcam_address,omitempty" json:"alamat"`
		NamaKabupaten string `bson:"korcam_city,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string `bson:"korcam_district,omitempty" json:"nama_kecamatan"`
		NamaKelurahan string `bson:"korcam_subdistrict,omitempty" json:"nama_kelurahan"`
		Jaringan      string `bson:"korcam_network,omitempty" json:"jaringan"`
	}

	PayloadStoreCoordinatorDistrict struct {
		KorcamName     string `json:"full_name"`
		KorcamNik      string `json:"nik"`
		KorcamPhone    string `json:"no_handphone"`
		KorcamAge      int    `json:"age"`
		KorcamGender   string `json:"gender"`
		KorcamAddress  string `json:"address"`
		KorcamCity     string `json:"city"`
		KorcamDistrict string `json:"district"`
		KorcamNetwork  string `json:"jaringan"`
	}

	PayloadUpdateCoordinatorDistrict struct {
		KorcamName     string `json:"full_name"`
		KorcamNik      string `json:"nik"`
		KorcamPhone    string `json:"no_handphone"`
		KorcamAge      int    `json:"age"`
		KorcamGender   string `json:"gender"`
		KorcamAddress  string `json:"address"`
		KorcamCity     string `json:"city"`
		KorcamDistrict string `json:"district"`
		KorcamNetwork  string `json:"jaringan"`
	}

	CoordinationDistrictFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Jaringan      string `json:"jaringan"`
		Nama          string `json:"nama"`
	}
)
