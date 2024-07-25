package repo

import "time"

type UserModel struct {
	UserID   string `db:"user_id"`
	Username string
	Password string
	Email    string
}

type InvoiceModel struct {
	InvoiceID       int     `db:"invoiceid"`
	IssueDate       []byte  `db:"issuedate"`
	Subject         string  `db:"subject"`
	TotalItems      int     `db:"totalitems"`
	CustomerID      int     `db:"customerid"`
	CustomerName    string  `db:"customername"`
	CustomerAddress string  `db:"customeraddress"`
	Amount          float64 `db:"amount"`
	DueDate         []byte  `db:"duedate"`
	Status          string  `db:"status"`
}

type InvoiceModelInsert struct {
	InvoiceID    int       `db:"invoiceid,omitempty"`
	IssueDate    time.Time `db:"issuedate"`
	Subject      string    `db:"subject"`
	TotalItems   int       `db:"totalitems"`
	CustomerID   int       `db:"customerid"`
	CustomerName string    `db:"customername"`
	Amount       float64   `db:"amount"`
	DueDate      time.Time `db:"duedate"`
}

type CustomerModel struct {
	CustomerID      int    `db:"customerid"`
	CustomerName    string `db:"customername"`
	CustomerAddress string `db:"customeraddress"`
}

type ItemModel struct {
	ItemID    int     `db:"itemid"`
	ItemName  string  `db:"itemname"`
	ItemType  string  `db:"itemtype"`
	UnitPrice float64 `db:"unitprice"`
}

type DetailItemModel struct {
	DetailItemID int     `db:"detailitemid,omitempty"`
	InvoiceID    int     `db:"invoiceid"`
	ItemID       int     `db:"itemid"`
	ItemName     string  `db:"itemname,omitempty"`
	Qty          int     `db:"qty"`
	Amount       float64 `db:"amount"`
}
