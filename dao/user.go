package dao

import (
	"GoPractice/entry"
	"GoPractice/models"
	"database/sql"
)

type User struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{DB: db}
}

func (dao *User) GetUserByID(id int64) (*models.Users, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := dao.DB.QueryRow(query, id)

	var user models.Users
	err := row.Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email,
		&user.CreatedBy, &user.CreatedOn, &user.LastModifiedOn, &user.Version,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *User) CreateUser(entry entry.UserEntry) (*models.Users, error) {
	query := `INSERT INTO users (first_name, last_name, email, created_by, created_on, last_modified_on, version)
		VALUES (?, ?, ?, ?)`
	result, err := dao.DB.Exec(query, entry.FirstName, entry.LastName, entry.Email)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return dao.GetUserByID(lastInsertId)
}
