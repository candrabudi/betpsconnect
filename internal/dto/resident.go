package dto

type (
	ResultTpsResidents struct {
		Items    []FindTpsResidents `json:"items"`
		Metadata MetaData           `json:"metadata"`
	}
	ResultResident struct {
		Items    []FindAllResident `json:"items"`
		Metadata MetaData          `json:"metadata"`
	}

	MetaData struct {
		TotalResults int `json:"total_results"`
		Limit        int `json:"limit"`
		Offset       int `json:"offset"`
		Count        int `json:"count"`
	}
	FindAllResident struct {
		ID             int    `json:"id"`
		Nama           string `json:"nama"`
		Alamat         string `json:"alamat"`
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
		Status         string `json:"status"`
		TanggalLahir   string `json:"tanggal_lahir"`
		StatusTpsLabel string `json:"status_tps_label"`
		Usia           int    `json:"usia"`
		Telp           string `json:"telp"`
		Tps            string `json:"tps"`
	}
	FindTpsResidents struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"nama,omitempty" json:"nama"`
		JenisKelamin  string `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin"`
		NamaKabupaten string `bson:"nama_kabupaten,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string `bson:"nama_kecamatan,omitempty" json:"nama_kecamatan"`
		Nik           string `bson:"nik,omitempty" json:"nik"`
		TanggalLahir  string `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir"`
		Usia          string `bson:"usia,omitempty" json:"usia"`
		Tps           string `bson:"tps,omitempty" json:"tps"`
		Status        string `bson:"status,omitempty" json:"status"`
		IsTrue        int    `bson:"is_true,omitempty" json:"is_true"`
		IsFalse       int    `bson:"is_false,omitempty" json:"is_false"`
	}

	DetailResident struct {
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

	ResidentFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		TPS           string `json:"tps"`
		Nama          string `json:"nama"`
	}

	PayloadStoreResident struct {
		Nama          string `json:"nama" binding:"required"`
		Alamat        string `json:"alamat" binding:"required"`
		JenisKelamin  string `json:"jenis_kelamin" binding:"required"`
		NamaKabupaten string `json:"nama_kabupaten" binding:"required"`
		NamaKecamatan string `json:"nama_kecamatan" binding:"required"`
		NamaKelurahan string `json:"nama_kelurahan" binding:"required"`
		Nik           string `json:"nik" binding:"required"`
		Rt            string `json:"rt" binding:"required"`
		Rw            string `json:"rw" binding:"required"`
		Usia          int    `json:"usia" binding:"required"`
		Telp          string `json:"Telp" binding:"required"`
		Tps           string `json:"tps" binding:"required"`
	}

	PayloadUpdateValidInvalid struct {
		ResidentID []int `json:"resident_id" binding:"required"`
		IsTrue     bool  `json:"is_true"`
	}

	ResultsDuplicateUpdateResident struct {
		Nama string
		NIK  string
	}

	FindAllResidentGrouped struct {
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Count         int32  `json:"count"`
	}

	KecamatanInKabupaten struct {
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Count         int32  `json:"count"`
	}

	FindTpsByDistrict struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		TotalTps      string `json:"total_tps"`
	}
)
