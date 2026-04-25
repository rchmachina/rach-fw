package request

type CreateAuth struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Name         string `json:"name"`
}
