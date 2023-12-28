package dto

type (
	GetByDistrict struct {
		ID            int    `json:"id"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKelurahan string `json:"nama_kelurahan"`
	}
)
