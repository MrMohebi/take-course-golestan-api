package faces

type BuyCodeRes struct {
	HasCode  bool   `json:"hasCode"`
	PayLink  string `json:"payLink"`
	DaysLeft int    `json:"daysLeft"`
}

type BuyCodeReq struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
	Email string `json:"email" form:"email"`
}
