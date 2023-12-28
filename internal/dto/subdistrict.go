package dto

type (
	GetByDistrict struct {
		ID            int    `json:"id"`
		NamaKabupaten string `json:"nama_kabupaten"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
	}
)
