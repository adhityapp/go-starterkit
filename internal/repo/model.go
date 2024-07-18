package repo

type UserModel struct {
	UserID   string `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type EmployeeNameModel struct {
	Firstname string
	Lastname  string
}

type EmployeeModel struct {
	EmployeeNameModel
	EmployeeID  string `db:"employee_id"`
	Salary      string
	SalaryNow   string `db:"salary_now"`
	ReviewCount string `db:"review_count"`
}
