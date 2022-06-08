package faces

type BuyCodeRes struct {
	Link string `json:"link"`
}

type BuyCodeReq struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
	Email string `json:"email" form:"email"`
}
