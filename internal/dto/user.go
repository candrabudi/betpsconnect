package dto

type (
	PayloadLogin struct {
		Email    string `json:"email", binding:"required"`
		Password string `json:"password", binding:"required"`
	}

	ResultLogin struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		Regency  string `json:"regency"`
		Status   string `json:"status"`
		Role     string `json:"role"`
		Token    string `json:"token"`
	}

	FindOneUser struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		Regency  string `json:"regency"`
		Status   string `json:"status"`
		Role     string `json:"role"`
	}
)
