package dto

type (
	GetByCity struct {
		ID            int    `json:"id"`
		NamaKecamatan string `json:"nama_kecamatan"`
		NamaKabupaten string `json:"nama_kabupaten"`
	}
)
