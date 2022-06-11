package faces

type PayVerifyReq struct {
	IdpayId         string `json:"id" form:"id" `
	Status          int    `json:"status,string" form:"status" `
	IdpayTrackingId int64  `json:"tracking_id,string" form:"tracking_id" `
	OrderId         string `json:"order_id" form:"order_id" `
	Amount          int64  `json:"amount,string" form:"amount" `
	CardNumber      string `json:"card_no" form:"card_no" `
	CardNumberHash  string `json:"hashed_card_no" form:"hashed_card_no" `
	PaidAt          int64  `json:"date,string" form:"date" `
}

type PayVerifyRes struct {
	SentSMS    bool   `json:"sentSMS"`
	ActiveCode string `json:"activeCode"`
}
