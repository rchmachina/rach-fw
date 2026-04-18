package request


type CreateUser struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedBy    *int64    `json:"created_by"`
}
