package repo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/adhityapp/go-starterkit/pkg/utils"
)

func (r RepoClient) GetInvoice(c context.Context) ([]InvoiceModel, error) {
	var im []InvoiceModel
	query := "SELECT invoiceid, issuedate, subject, totalitems, customerid, amount, duedate, status FROM invoice;"
	rows, err := r.DB.MySQLDB().QueryxContext(c, query)
	if err != nil {
		utils.LogError("repo", "select-db", err)
		return nil, err
	}
	for rows.Next() {
		var m InvoiceModel
		err = rows.StructScan(&m)
		if err != nil {
			utils.LogError("repo", "struct-scan", err)
			return nil, err
		}
		im = append(im, m)
	}
	return im, nil
}

func (r RepoClient) GetInvoiceByID(c context.Context, id int) (*InvoiceModel, error) {
	var im InvoiceModel
	query := `SELECT 
    i.invoiceid, 
    i.issuedate, 
    i.subject, 
    i.totalitems, 
    i.customerid, 
    i.amount, 
    i.duedate, 
    i.status,
    c.customername,
	c.customeraddress
FROM 
    invoice i
INNER JOIN 
    customer c ON i.customerid = c.customerid where invoiceid = ?`
	err := r.DB.MySQLDB().QueryRowxContext(c, query, id).StructScan(&im)
	if err != nil {
		utils.LogError("repo-get-invoice-by-id", "select-db", err)
		return nil, err
	}
	return &im, nil
}

func (r RepoClient) GetInvoiceDetail(c context.Context, id int) ([]DetailItemModel, error) {
	var im []DetailItemModel
	query := `SELECT 
    d.detailitemid, 
    d.invoiceid, 
    d.itemid, 
    d.qty, 
    d.amount,
    i.itemname
FROM 
    detail_items d
INNER JOIN 
    items i ON d.itemid = i.itemid WHERE invoiceid = ?`
	rows, err := r.DB.MySQLDB().QueryxContext(c, query, id)
	if err != nil {
		utils.LogError("repo-get-invoice-detail", "select-db", err)
		return nil, err
	}
	for rows.Next() {
		var m DetailItemModel
		err = rows.StructScan(&m)
		if err != nil {
			utils.LogError("repo", "struct-scan", err)
			return nil, err
		}
		im = append(im, m)
	}
	return im, nil
}

func (r RepoClient) GetCustomer(c context.Context) ([]CustomerModel, error) {
	var im []CustomerModel
	query := `SELECT customerid, customername, customeraddress from customer`
	rows, err := r.DB.MySQLDB().QueryxContext(c, query)
	if err != nil {
		utils.LogError("repo-get-invoice-detail", "select-db", err)
		return nil, err
	}
	for rows.Next() {
		var m CustomerModel
		err = rows.StructScan(&m)
		if err != nil {
			utils.LogError("repo", "struct-scan", err)
			return nil, err
		}
		im = append(im, m)
	}
	return im, nil
}

func (r RepoClient) GetItem(c context.Context) ([]ItemModel, error) {
	var im []ItemModel
	query := `SELECT itemid, itemname, itemtype, unitprice from items`
	rows, err := r.DB.MySQLDB().QueryxContext(c, query)
	if err != nil {
		utils.LogError("repo-get-invoice-detail", "select-db", err)
		return nil, err
	}
	for rows.Next() {
		var m ItemModel
		err = rows.StructScan(&m)
		if err != nil {
			utils.LogError("repo", "struct-scan", err)
			return nil, err
		}
		im = append(im, m)
	}
	return im, nil
}

func (r RepoClient) DeleteDetailItem(c context.Context, id int) error {
	query := `delete from detail_items where detailitemid = ?`
	_, err := r.DB.MySQLDB().ExecContext(c, query, id)
	if err != nil {
		utils.LogError("repo-get-invoice-detail", "select-db", err)
		return err
	}
	return nil
}

func (r RepoClient) AddInvoice(c context.Context, im InvoiceModelInsert) (*int, error) {
	query := `INSERT INTO invoice (issuedate, subject, totalitems, customerid, amount, duedate)
	VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.DB.MySQLDB().ExecContext(c, query,
		im.IssueDate,
		im.Subject,
		im.TotalItems,
		im.CustomerID,
		im.Amount,
		im.DueDate,
	)
	if err != nil {
		utils.LogError("repo-add-invoice", "insert", err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	var ids int = int(id)
	return &ids, nil
}

func (r RepoClient) AddInvoiceDetail(c context.Context, dm []DetailItemModel) error {
	query := `
        INSERT INTO detail_items (invoiceid, itemid, qty, amount)
        VALUES `

	valueStrings := make([]string, len(dm))
	valueArgs := make([]interface{}, len(dm)*4)

	for i, item := range dm {
		valueStrings[i] = "(?, ?, ?, ?)"
		valueArgs[i*4] = item.InvoiceID
		valueArgs[i*4+1] = item.ItemID
		valueArgs[i*4+2] = item.Qty
		valueArgs[i*4+3] = item.Amount
	}

	query += strings.Join(valueStrings, ",")

	_, err := r.DB.MySQLDB().ExecContext(c, query, valueArgs...)
	if err != nil {
		utils.LogError("repo-add-invoice-detail", "select-db", err)
		return err
	}
	return nil
}

func (r RepoClient) UpdateInvoice(c context.Context, im InvoiceModelInsert) error {
	query := `UPDATE invoice
	SET issuedate = ?, subject = ?, totalitems = ?, customerid = ?, amount = ?, duedate = ?
	WHERE invoiceid = ?`

	_, err := r.DB.MySQLDB().ExecContext(c, query,
		im.IssueDate,
		im.Subject,
		im.TotalItems,
		im.CustomerID,
		im.Amount,
		im.DueDate,
		im.InvoiceID,
	)
	if err != nil {
		utils.LogError("repo-add-invoice", "insert", err)
		return err
	}
	return nil
}

func (r RepoClient) UpdateInvoiceItem(c context.Context, items []DetailItemModel) error {

	var invoiceIdCaseStmt, itemIdCaseStmt, qtyCaseStmt, amountCaseStmt strings.Builder
	ids := []string{}

	for _, item := range items {
		ids = append(ids, strconv.Itoa(item.DetailItemID))
		invoiceIdCaseStmt.WriteString(fmt.Sprintf(`
            WHEN %d THEN %d
        `, item.DetailItemID, item.InvoiceID))
		itemIdCaseStmt.WriteString(fmt.Sprintf(`
            WHEN %d THEN %d
        `, item.DetailItemID, item.ItemID))
		qtyCaseStmt.WriteString(fmt.Sprintf(`
            WHEN %d THEN %d
        `, item.DetailItemID, item.Qty))
		amountCaseStmt.WriteString(fmt.Sprintf(`
            WHEN %d THEN %f
        `, item.DetailItemID, item.Amount))
	}

	query := fmt.Sprintf(`
	UPDATE detail_items
	SET
		invoiceid = CASE detailitemid
		%s
		END,
		itemid = CASE detailitemid
		%s
		END,
		qty = CASE detailitemid
		%s
		END,
		amount = CASE detailitemid
		%s
		END
	WHERE detailitemid IN (%s);
`,
		invoiceIdCaseStmt.String(),
		itemIdCaseStmt.String(),
		qtyCaseStmt.String(),
		amountCaseStmt.String(),
		strings.Join(ids, ","))

	_, err := r.DB.MySQLDB().ExecContext(c, query)
	if err != nil {
		utils.LogError("repo-update-invoice-item", "insert", err)
		return err
	}
	return nil
}
