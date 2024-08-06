package viewmodel

type UserViewModel struct {
	UserID   string `json:"UserID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Role     string `json:"Role"`
}
