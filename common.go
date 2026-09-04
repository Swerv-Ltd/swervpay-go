package swervpay

type PageAndLimitQuery struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Email string `json:"email,omitempty"`
}
