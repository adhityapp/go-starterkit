package repo

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (r RepoClient) GetSmithActiveEmployee(c context.Context) ([]EmployeeNameModel, error) {
	var um []EmployeeNameModel
	query := `select firstname,lastname from employee WHERE lastname LIKE 'Smith%' AND terminationdate IS NULL ORDER BY lastname, firstname`
	rows, err := r.DB.Dbr().QueryxContext(c, query)
	if err != nil {
		logrus.
			WithField("repo", "getsmith").
			Error(err)
		return nil, err
	}
	for rows.Next() {
		var m EmployeeNameModel
		rows.StructScan(&m)
		um = append(um, m)
	}
	return um, nil
}

func (r RepoClient) GetEmployeeNoReview(c context.Context) ([]EmployeeNameModel, error) {
	var um []EmployeeNameModel
	query := `SELECT e.firstname, e.lastname
	FROM employee e 
	LEFT JOIN annualreview ar 
	ON e.employee_id = ar.employee_id 
	WHERE ar.anid IS NULL 
	ORDER BY e.hiredate;`

	rows, err := r.DB.Dbr().QueryxContext(c, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m EmployeeNameModel
		rows.StructScan(&m)
		um = append(um, m)
	}
	return um, nil
}

func (r RepoClient) GetEmployeeDifferentDay(c context.Context) (*int, error) {
	var daydifferent int
	query := `SELECT DATE_PART('day', (SELECT hiredate FROM employee WHERE terminationdate IS NULL ORDER BY hiredate DESC LIMIT 1) - (SELECT hiredate FROM employee WHERE terminationdate IS NULL ORDER BY hiredate ASC LIMIT 1)) AS difference_in_days;`
	err := r.DB.Dbr().QueryRow(query).Scan(&daydifferent)
	if err != nil {
		return nil, err
	}
	return &daydifferent, nil
}

func (r RepoClient) GetSalary(c context.Context) ([]EmployeeModel, error) {
	var em []EmployeeModel
	query := `WITH SalaryIncrease AS (
		SELECT
			e.employee_id,
			e.firstname,
			e.lastname,
			e.salary::FLOAT,
			e.salary::FLOAT * POWER(1.15, 2024 - EXTRACT(YEAR FROM e.hiredate)) AS salary_now,
			COUNT(ar.anid) AS review_count
		FROM
			employee e
		LEFT JOIN
			annualreview ar ON e.employee_id = ar.employee_id
		WHERE
			e.hiredate <= '2024-12-31'
		GROUP BY
			e.employee_id
	)
	SELECT
		firstname,
		lastname,
		employee_id,
		salary,
		salary_now,
		review_count
	FROM
		SalaryIncrease
	ORDER BY
		salary_now DESC,
		review_count ASC;`

	rows, err := r.DB.Dbr().QueryxContext(c, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m EmployeeModel
		err = rows.StructScan(&m)
		if err != nil {
			return nil, err
		}
		em = append(em, m)
	}
	return em, nil
}
