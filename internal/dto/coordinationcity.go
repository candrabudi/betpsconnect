package dto

type (
	ResultAllCoordinatorCity struct {
		Items    []FindCoordinatorCity `json:"items"`
		Metadata MetaData              `json:"metadata"`
	}

	FindCoordinatorCity struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"coordinator_name,omitempty" json:"nama"`
		Nik           string `bson:"coordinator_nik,omitempty" json:"nik"`
		Usia          int    `bson:"coordinator_age,omitempty" json:"usia"`
		Alamat        string `bson:"coordinator_address,omitempty" json:"alamat"`
		NamaKabupaten string `bson:"coordinator_city,omitempty" json:"nama_kabupaten"`
	}

	searchCoordination struct {
		Nama string `json:"nama"`
	}
)
