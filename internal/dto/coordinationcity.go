package dto

import "time"

type (
	ResultAllCoordinatorCity struct {
		Items    []FindCoordinatorCity `json:"items"`
		Metadata MetaData              `json:"metadata"`
	}

	FindCoordinatorCity struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"korkab_name,omitempty" json:"nama"`
		NoHandphone   string `bson:"korkab_phone,omitempty" json:"telp"`
		Nik           string `bson:"korkab_nik,omitempty" json:"nik"`
		Usia          int    `bson:"korkab_age,omitempty" json:"usia"`
		Gender        string `bson:"korkab_gender,omitempty" json:"jenis_kelamin"`
		Alamat        string `bson:"korkab_address,omitempty" json:"alamat"`
		NamaKabupaten string `bson:"korkab_city,omitempty" json:"nama_kabupaten"`
		Jaringan      string `bson:"korkab_network,omitempty" json:"jaringan"`
	}

	ExportCoordinatorCity struct {
		ID            int       `bson:"id,omitempty" json:"id"`
		Nama          string    `bson:"korkab_name,omitempty" json:"nama"`
		NoHandphone   string    `bson:"korkab_phone,omitempty" json:"telp"`
		Nik           string    `bson:"korkab_nik,omitempty" json:"nik"`
		Usia          int       `bson:"korkab_age,omitempty" json:"usia"`
		Gender        string    `bson:"korkab_gender,omitempty" json:"jenis_kelamin"`
		Alamat        string    `bson:"korkab_address,omitempty" json:"alamat"`
		NamaKabupaten string    `bson:"korkab_city,omitempty" json:"nama_kabupaten"`
		Jaringan      string    `bson:"korkab_network,omitempty" json:"jaringan"`
		CreatedAt     time.Time `bson:"created_at,omitempty" json:"created_at"`
		UpdatedAt     time.Time `bson:"updated_at,omitempty" json:"updated_at"`
	}

	CoordinationCityFilter struct {
		Nama       string `json:"nama"`
		KorkabCity string `bson:"korkab_city,omitempty" json:"nama_kabupaten"`
		Jaringan   string `json:"jaringan"`
	}

	PayloadStoreCoordinatorCity struct {
		KorkabName    string `json:"full_name"`
		KorkabNik     string `json:"nik"`
		KorkabPhone   string `json:"no_handphone"`
		KorkabAge     int    `json:"age"`
		KorkabGender  string `json:"Gender"`
		KorkabAddress string `json:"address"`
		KorkabCity    string `json:"city"`
		KorkabNetwork string `json:"jaringan"`
	}
)
