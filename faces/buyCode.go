package faces

type BuyCodeRes struct {
	HasCode  bool   `json:"hasCode"`
	PayLink  string `json:"payLink,omitempty"`
	DaysLeft int    `json:"daysLeft,omitempty"`
}

type BuyCodeReq struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
	Email string `json:"email" form:"email"`
}
