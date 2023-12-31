package dto

type (
	TrueResidentPayload struct {
		FullName    string `json:"full_name"`
		NIK         string `json:"nik"`
		Gender      string `json:"gender"`
		District    string `json:"district"`
		Subdistrict string `json:"subdistrict"`
		City        string `json:"city"`
		BirthDate   string `json:"birth_date"`
		BirthPlace  string `json:"birth_place"`
		TPS         string `json:"tps"`
	}
)
