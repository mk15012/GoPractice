package models

import "time"

type Users struct {
	ID             int64     `db:"id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Email          string    `db:"email"`
	CreatedBy      string    `db:"created_by"`
	CreatedOn      time.Time `db:"created_on"`
	LastModifiedOn time.Time `db:"last_modified_on"`
	Version        int       `db:"version"`
}
