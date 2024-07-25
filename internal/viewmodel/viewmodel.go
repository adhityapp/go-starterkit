package viewmodel

import "time"

type RequestInvoice struct {
	InvoiceID  int             `json:"invoice_id,omitempty"`
	Subject    string          `json:"subject"`
	IssueDate  time.Time       `json:"issue_date"`
	DueDate    time.Time       `json:"due_date"`
	CustomerID int             `json:"customer_id"`
	TotalItems int             `json:"total_items"`
	Amount     float64         `json:"amount"`
	Detail     []RequestDetail `json:"detail"`
}

type RequestDetail struct {
	DetailItemID int     `json:"detail_item_id,omitempty"`
	ItemID       int     `json:"item_id"`
	Qty          int     `json:"qty"`
	Amount       float64 `json:"amount"`
}
