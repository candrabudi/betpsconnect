package dto

type (
	ResultAllCoordinatorCity struct {
		Items    []FindCoordinatorCity `json:"items"`
		Metadata MetaData              `json:"metadata"`
	}

	FindCoordinatorCity struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"korkab_name,omitempty" json:"name"`
		NoHandphone   string `bson:"korkab_phone,omitempty" json:"no_handphone"`
		Nik           string `bson:"korkab_nik,omitempty" json:"nik"`
		Usia          int    `bson:"korkab_age,omitempty" json:"age"`
		Alamat        string `bson:"korkab_address,omitempty" json:"address"`
		NamaKabupaten string `bson:"korkab_city,omitempty" json:"city"`
	}

	searchCoordination struct {
		Nama string `json:"nama"`
	}

	PayloadStoreCoordinatorCity struct {
		KorkabName    string `json:"korkab_name"`
		KorkabNik     string `json:"korkab_nik"`
		KorkabPhone   string `json:"korkab_phone"`
		KorkabAge     string `json:"korkab_age"`
		KorkabAddress string `json:"korkab_address"`
		KorkabCity    string `json:"korkab_city"`
	}
)
