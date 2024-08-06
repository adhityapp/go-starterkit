package repo

type UserModel struct {
	UserID   string `db:"user_id"`
	Username string
	Password string
	Email    string
	Role     string `db:"user_role"`
}
