package faces

type LoginRes struct {
	Token string `json:"token"`
}

type LoginReq struct {
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Email    string `json:"email" form:"email"`
}
