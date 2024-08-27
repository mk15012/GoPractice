package models

import "time"

type Products struct {
	ID             int64     `db:"id"`
	ProductName    string    `db:"product_name"`
	Description    string    `db:"description"`
	Quantity       string    `db:"quantity"`
	CreatedBy      string    `db:"created_by"`
	CreatedOn      time.Time `db:"created_on"`
	LastModifiedOn time.Time `db:"last_modified_on"`
	Version        int       `db:"version"`
}
