package model

type GetLoginUser struct {
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password_hash"`
	Id       string `gorm:"column:id"`
	Name     string `json:"column:name"`
}

type GetToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
