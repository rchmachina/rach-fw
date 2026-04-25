package request

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	RefreshToken string `json:"refreshToken"`
}
