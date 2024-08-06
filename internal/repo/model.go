package repo

type UserModel struct {
	UserID   string `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Role     string `db:"user_role"`
}
