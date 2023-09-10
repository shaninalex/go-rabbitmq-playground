package models

type NewUser struct {
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	FirstName    string `json:"lastname" binding:"required"`
	LastName     string `json:"firstname" binding:"required"`
	PasswordHash string `json:"_" binding:"required"`
}

func (a *NewUser) Create() error {

	return nil
}
