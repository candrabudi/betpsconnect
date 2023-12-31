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

	DetailTrueResident struct {
		ID             int    `json:"id"`
		Nama           string `json:"nama"`
		Alamat         string `json:"alamat"`
		Difabel        string `json:"difabel"`
		Ektp           string `json:"ektp"`
		Email          string `json:"email"`
		JenisKelamin   string `json:"jenis_kelamin"`
		Kawin          string `json:"kawin"`
		NamaKabupaten  string `json:"nama_kabupaten"`
		NamaKecamatan  string `json:"nama_kecamatan"`
		NamaKelurahan  string `json:"nama_kelurahan"`
		Nik            string `json:"nik"`
		Nkk            string `json:"nkk"`
		NoKtp          string `json:"no_ktp"`
		Rt             string `json:"rt"`
		Rw             string `json:"rw"`
		SaringanID     string `json:"saringan_id"`
		Status         string `json:"status"`
		StatusTpsLabel string `json:"status_tps_label"`
		TanggalLahir   string `json:"tanggal_lahir"`
		Usia           int    `json:"usia"`
		TempatLahir    string `json:"tampat_lahir"`
		Telp           string `json:"telp"`
		Tps            string `json:"tps"`
		IsTrue         int    `json:"is_true"`
		IsFalse        int    `json:"is_false"`
	}
)
