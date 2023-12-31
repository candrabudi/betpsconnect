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
)
