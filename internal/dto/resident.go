package dto

type (
	ResultTpsResidents struct {
		Items    []FindTpsResidents `json:"items"`
		Metadata MetaData           `json:"metadata"`
	}

	ResultValidateResidents struct {
		Items    []FindValidateResidents `json:"items"`
		Metadata MetaData                `json:"metadata"`
	}

	MetaData struct {
		TotalResults int `json:"total_results"`
		Limit        int `json:"limit"`
		Offset       int `json:"offset"`
		Count        int `json:"count"`
	}

	FindTpsResidents struct {
		ID             int    `bson:"id,omitempty" json:"id"`
		Nama           string `bson:"nama,omitempty" json:"nama"`
		JenisKelamin   string `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin"`
		NamaKabupaten  string `bson:"nama_kabupaten,omitempty" json:"nama_kabupaten"`
		NamaKecamatan  string `bson:"nama_kecamatan,omitempty" json:"nama_kecamatan"`
		Nik            string `bson:"nik,omitempty" json:"nik"`
		TanggalLahir   string `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir"`
		Usia           int    `bson:"usia,omitempty" json:"usia"`
		TPS            string `bson:"tps,omitempty" json:"tps"`
		Status         string `bson:"status,omitempty" json:"status"`
		IsVerification int    `bson:"is_verification,omitempty" json:"is_verification"`
		StatusTPSLabel string `bson:"status_tps_label,omitempty" json:"status_tps_label"`
		TempatLahir    string `bson:"tempat_lahir,omitempty" json:"tempat_lahir"`
		Telp           string `bson:"telp,omitempty" json:"telp"`
		NoKTP          string `bson:"no_ktp,omitempty" json:"no_ktp"`
		Nkk            string `bson:"nkk,omitempty" json:"nkk"`
		Difabel        string `bson:"difabel,omitempty" json:"difabel"`
		Kawin          string `bson:"kawin,omitempty" json:"kawin"`
		RT             string `bson:"rt,omitempty" json:"rt"`
		RW             string `bson:"rw,omitempty" json:"rw"`
		Alamat         string `bson:"alamat,omitempty" json:"alamat"`
	}

	FindValidateResidents struct {
		ID             int    `bson:"id,omitempty" json:"id"`
		Nama           string `bson:"nama,omitempty" json:"nama"`
		JenisKelamin   string `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin"`
		NamaKabupaten  string `bson:"nama_kabupaten,omitempty" json:"nama_kabupaten"`
		NamaKecamatan  string `bson:"nama_kecamatan,omitempty" json:"nama_kecamatan"`
		Nik            string `bson:"nik,omitempty" json:"nik"`
		TanggalLahir   string `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir"`
		Usia           int    `bson:"usia,omitempty" json:"usia"`
		TPS            string `bson:"tps,omitempty" json:"tps"`
		Status         string `bson:"status,omitempty" json:"status"`
		IsVerification int    `bson:"is_verification,omitempty" json:"is_verification"`
		StatusTPSLabel string `bson:"status_tps_label,omitempty" json:"status_tps_label"`
		TempatLahir    string `bson:"tempat_lahir,omitempty" json:"tempat_lahir"`
		Telp           string `bson:"telp,omitempty" json:"telp"`
		NoKTP          string `bson:"no_ktp,omitempty" json:"no_ktp"`
		Nkk            string `bson:"nkk,omitempty" json:"nkk"`
		Difabel        string `bson:"difabel,omitempty" json:"difabel"`
		Kawin          string `bson:"kawin,omitempty" json:"kawin"`
		RT             string `bson:"rt,omitempty" json:"rt"`
		RW             string `bson:"rw,omitempty" json:"rw"`
		Alamat         string `bson:"alamat,omitempty" json:"alamat"`
	}

	DetailResident struct {
		ID             int    `json:"id" bson:"id"`
		Nama           string `json:"nama" bson:"nama"`
		Alamat         string `json:"alamat" bson:"alamat"`
		Difabel        string `json:"difabel" bson:"difabel"`
		Ektp           string `json:"ektp" bson:"ektp"`
		Email          string `json:"email" bson:"email"`
		JenisKelamin   string `json:"jenis_kelamin" bson:"jenis_kelamin"`
		Kawin          string `json:"kawin" bson:"kawin"`
		NamaKabupaten  string `json:"nama_kabupaten" bson:"nama_kabupaten"`
		NamaKecamatan  string `json:"nama_kecamatan" bson:"nama_kecamatan"`
		NamaKelurahan  string `json:"nama_kelurahan" bson:"nama_kelurahan"`
		Nik            string `json:"nik" bson:"nik"`
		Nkk            string `json:"nkk" bson:"nkk"`
		NoKtp          string `json:"no_ktp" bson:"no_ktp"`
		Rt             string `json:"rt" bson:"rt"`
		Rw             string `json:"rw" bson:"rw"`
		SaringanID     string `json:"saringan_id" bson:"saringan_id"`
		Status         string `json:"status" bson:"status"`
		StatusTpsLabel string `json:"status_tps_label" bson:"status_tps_label"`
		TanggalLahir   string `json:"tanggal_lahir" bson:"tanggal_lahir"`
		Usia           int    `json:"usia" bson:"usia"`
		TempatLahir    string `json:"tempat_lahir" bson:"tempat_lahir"`
		Telp           string `json:"telp" bson:"telp"`
		Tps            string `json:"tps" bson:"tps"`
		IsVerification int    `json:"is_verification" bson:"is_verification"`
	}

	ResidentFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		TPS           string `json:"tps"`
		Nama          string `json:"nama"`
	}

	PayloadStoreResident struct {
		FullName    string `json:"full_name" validate:"required"`
		Nik         string `json:"nik" validate:"required"`
		Gender      string `json:"gender" validate:"required"`
		District    string `json:"district" validate:"required"`
		Subdistrict string `json:"subdistrict" validate:"required"`
		City        string `json:"city" validate:"required"`
		Address     string `json:"address" validate:"required"`
		Age         int    `json:"age" validate:"required"`
		NoHandphone string `json:"no_handphone" validate:"required"`
		TPS         string `json:"tps" validate:"required"`
	}

	PayloadUpdateValidInvalid struct {
		IsTrue bool                          `json:"is_true", binding:"required"`
		Items  []PayloadBulkValidateResident `json:"items", binding:"required"`
	}

	PayloadBulkValidateResident struct {
		ID          int    `json:"id", binding:"required"`
		NoHandphone string `json:"no_handphone", binding:"required"`
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
