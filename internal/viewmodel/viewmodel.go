package viewmodel

type UserViewModel struct {
	UserID   string `json:"UserID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type EmployeeNameViewModel struct {
	Firstname string `json:"FirstName"`
	Lastname  string `json:"LastName"`
}

type EmployeeViewModel struct {
	EmployeeNameViewModel
	EmployeeID  string `json:"EmployeeID"`
	Salary      string `json:"Salary"`
	SalaryNow   string `json:"SalaryNow"`
	ReviewCount string `json:"ReviewCount"`
}
