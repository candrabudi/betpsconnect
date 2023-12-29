package dto

type (
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
		StatusTpsLabel string `json:"status_tps_label"`
		Usia           int    `json:"usia"`
		Telp           string `json:"telp"`
		Tps            string `json:"tps"`
	}

	DetailResident struct {
		ID                 int    `json:"id"`
		Nama               string `json:"nama"`
		Alamat             string `json:"alamat"`
		Difabel            string `json:"difabel"`
		Ektp               string `json:"ektp"`
		Email              string `json:"email"`
		JenisKelamin       string `json:"jenisKelamin"`
		Kawin              string `json:"kawin"`
		KabID              int    `json:"KabID"`
		KecID              int    `json:"KecID"`
		KelID              int    `json:"KelID"`
		NamaKabupaten      string `json:"namakabupaten"`
		NamaKecamatan      string `json:"namaKecamatan"`
		NamaKelurahan      string `json:"namaKelurahan"`
		Nik                string `json:"nik"`
		Nkk                string `json:"nkk"`
		NoKtp              string `json:"noKtp"`
		Rt                 string `json:"rt"`
		Rw                 string `json:"rw"`
		SaringanID         string `json:"saringanId"`
		Status             string `json:"status"`
		StatusTpsLabel     string `json:"statusTpsLabel"`
		TanggalLahir       string `json:"tanggalLahir"`
		TempatLahir        string `json:"tempatLahir"`
		TanggalLahirString string `json:"tanggalLahirString"`
		Telp               string `json:"Telp"`
		Tps                string `json:"tps"`
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
