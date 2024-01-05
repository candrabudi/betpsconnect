package dto

type (
	ResultAllCoordinatorDistrict struct {
		Items    []FindCoordinatorCity `json:"items"`
		Metadata MetaData              `json:"metadata"`
	}

	FindCoordinatorDistrict struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"coordinator_name,omitempty" json:"name"`
		NoHandphone   string `bson:"coordinator_phone,omitempty" json:"no_handphone"`
		Nik           string `bson:"coordinator_nik,omitempty" json:"nik"`
		Usia          int    `bson:"coordinator_age,omitempty" json:"age"`
		Alamat        string `bson:"coordinator_address,omitempty" json:"address"`
		NamaKabupaten string `bson:"coordinator_city,omitempty" json:"city"`
	}

	PayloadStoreCoordinatorDistrict struct {
		KorkabID       int    `json:"korkab_id"`
		KorcamName     string `json:"korcam_name"`
		KorcamNik      string `json:"korcam_nik"`
		KorcamPhone    string `json:"korcam_phone"`
		KorcamAge      string `json:"korcam_age"`
		KorcamAddress  string `json:"korcam_address"`
		KorcamCity     string `json:"korcam_city"`
		KorcamDistrict string `json:"korcam_district"`
	}

	PayloadUpdateCoordinatorDistrict struct {
		KorcamName     string `json:"korcam_name"`
		KorcamNik      string `json:"korcam_nik"`
		KorcamPhone    string `json:"korcam_phone"`
		KorcamAge      string `json:"korcam_age"`
		KorcamAddress  string `json:"korcam_address"`
		KorcamCity     string `json:"korcam_city"`
		KorcamDistrict string `json:"korcam_district"`
	}
)
