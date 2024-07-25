package service

import "time"

type UserViewModel struct {
	UserID   string `json:"UserID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type InvoiceViewModel struct {
	InvoiceID  int       `json:"invoice_id"`
	IssueDate  time.Time `json:"issue_date"`
	Subject    string    `json:"subject"`
	TotalItems int       `json:"total_items"`
	CustomerID int       `json:"customer_id"`
	Amount     float64   `json:"amount"`
	DueDate    time.Time `json:"due_date"`
	Status     string    `json:"status"`
}

type CustomerViewModel struct {
	CustomerID      int    `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	CustomerAddress string `json:"customer_address"`
}

type ItemViewModel struct {
	ItemID    int     `json:"item_id"`
	ItemName  string  `json:"item_name"`
	ItemType  string  `json:"item_type"`
	UnitPrice float64 `json:"unit_price"`
}

type DetailItemViewModel struct {
	DetailItemID int     `json:"detail_item_id"`
	InvoiceID    int     `json:"invoice_id"`
	ItemID       int     `json:"item_id"`
	Qty          int     `json:"qty"`
	Amount       float64 `json:"amount"`
}

type InvoiceDetailViewModel struct {
	InvoiceID    int                    `json:"invoice_id"`
	IssueDate    time.Time              `json:"issue_date"`
	Subject      string                 `json:"subject"`
	TotalItems   int                    `json:"total_items"`
	CustomerID   int                    `json:"customer_id"`
	CustomerName string                 `json:"customer_name,omitempty"`
	CustomerAddr string                 `json:"customer_address,omitempty"`
	Amount       float64                `json:"amount"`
	DueDate      time.Time              `json:"due_date"`
	Status       string                 `json:"status"`
	Detail       []DetailItemViewModel2 `json:"detail"`
}

type DetailItemViewModel2 struct {
	DetailItemID int     `json:"detail_item_id"`
	ItemID       int     `json:"item_id"`
	ItemName     string  `json:"item_name,omitempty"`
	Qty          int     `json:"qty"`
	Amount       float64 `json:"amount"`
}

type DropdownViewModel struct {
	Customer []CustomerViewModel `json:"customer"`
	Items    []ItemViewModel     `json:"items"`
}
