package faces

type IsActiveCodeValidReq struct {
	Code string `json:"code" form:"code" validate:"required"`
}

type IsActiveCodeValidRes struct {
	IsValid  bool `json:"isValid"`
	DaysLeft int  `json:"daysLeft"`
}
