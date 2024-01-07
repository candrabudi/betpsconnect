package dto

import "time"

type (
	TrueResidentPayload struct {
		FullName    string `json:"full_name" validate:"required"`
		Address     string `json:"address" validate:"required"`
		Nik         string `json:"nik" validate:"required"`
		Gender      string `json:"gender" validate:"required"`
		District    string `json:"district" validate:"required"`
		Subdistrict string `json:"subdistrict" validate:"required"`
		City        string `json:"city" validate:"required"`
		Age         int    `json:"age" validate:"required"`
		NoHandphone string `json:"no_handphone" validate:"required"`
		TPS         string `json:"tps" validate:"required"`
		Jaringan    string `json:"jaringan"`
	}

	ResultAllTrueResident struct {
		Items    []FindTrueAllResident `json:"items"`
		Metadata MetaData              `json:"metadata"`
	}

	TrueResidentFilter struct {
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
		Jaringan      string `json:"jaringan"`
		TPS           string `json:"tps"`
		Nama          string `json:"nama"`
		IsManual      string `json:"is_manual"`
	}

	PayloadUpdateTrueResident struct {
		FullName    string `json:"full_name"`
		Address     string `json:"address"`
		Nik         string `json:"nik"`
		Gender      string `json:"gender"`
		District    string `json:"district"`
		Subdistrict string `json:"subdistrict"`
		City        string `json:"city"`
		Age         int    `json:"age"`
		NoHandphone string `json:"no_handphone"`
		TPS         string `json:"tps"`
	}

	FindTrueAllResident struct {
		ID            int    `bson:"id,omitempty" json:"id"`
		Nama          string `bson:"full_name,omitempty" json:"nama"`
		JenisKelamin  string `bson:"gender,omitempty" json:"jenis_kelamin"`
		NamaKabupaten string `bson:"city,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string `bson:"district,omitempty" json:"nama_kecamatan"`
		NamaKelurahan string `bson:"subdistrict,omitempty" json:"nama_kelurahan"`
		Address       string `bson:"address,omitempty" json:"alamat"`
		Nik           string `bson:"nik,omitempty" json:"nik"`
		Usia          int    `bson:"age,omitempty" json:"usia"`
		Tps           string `bson:"tps,omitempty" json:"tps"`
		NoHandphone   string `bson:"no_handphone,omitempty" json:"telp"`
		Jaringan      string `bson:"network,empty" json:"jaringan"`
		IsManual      int    `bson:"is_manual,omitempty" json:"is_manual"`
	}

	ExportTrueAllResident struct {
		Nama          string    `bson:"full_name,omitempty" json:"nama"`
		JenisKelamin  string    `bson:"gender,omitempty" json:"jenis_kelamin"`
		NamaKabupaten string    `bson:"city,omitempty" json:"nama_kabupaten"`
		NamaKecamatan string    `bson:"district,omitempty" json:"nama_kecamatan"`
		NamaKelurahan string    `bson:"subdistrict,omitempty" json:"nama_kelurahan"`
		Address       string    `bson:"address,omitempty" json:"alamat"`
		Nik           string    `bson:"nik,omitempty" json:"nik"`
		Usia          int       `bson:"age,omitempty" json:"usia"`
		Tps           string    `bson:"tps,omitempty" json:"tps"`
		NoHandphone   string    `bson:"no_handphone,omitempty" json:"telp"`
		Jaringan      string    `bson:"network,empty" json:"jaringan"`
		CreatedAt     time.Time `bson:"created_at,omitempty" json:"created_at"`
		UpdatedAt     time.Time `bson:"updated_at,omitempty" json:"updated_at"`
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
