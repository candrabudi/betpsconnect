package dto

type (
	TrueResidentPayload struct {
		FullName    string `json:"full_name"`
		Nik         string `json:"nik"`
		Gender      string `json:"gender"`
		District    string `json:"district"`
		Subdistrict string `json:"subdistrict"`
		City        string `json:"city"`
		Age         int    `json:"age"`
		NoHandphone string `json:"no_handphone"`
		TPS         string `json:"tps"`
	}

	ResultAllTrueResident struct {
		Items    []FindTrueAllResident `json:"items"`
		Metadata MetaData              `json:"metadata"`
	}

	FindTrueAllResident struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"full_name,omitempty" json:"nama"`
		JenisKelamin  string `bson:"gender,omitempty" json:"jenis_kelamin"`
		NamaKecamatan string `bson:"district,omitempty" json:"nama_kecamatan"`
		Nik           string `bson:"nik,omitempty" json:"nik"`
		TanggalLahir  string `bson:"birth_date,omitempty" json:"tanggal_lahir"`
		Usia          int    `bson:"age,omitempty" json:"usia"`
		Tps           string `bson:"tps,omitempty" json:"tps"`
	}
)
