package database

import (
	"GoPractice/entry"
	"GoPractice/models"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Product struct {
	DB  *sql.DB
	Log *logrus.Logger
}

func NewProduct(db *sql.DB) *Product {
	return &Product{DB: db}
}

func (dao *Product) GetProductByID(id int64) (*models.Products, error) {
	query := "SELECT * FROM products WHERE id = ?"
	row := dao.DB.QueryRow(query, id)

	var product models.Products
	err := row.Scan(
		&product.ID, &product.Description, &product.Quantity,
		&product.CreatedBy, &product.CreatedOn, &product.LastModifiedOn, &product.Version,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (dao *Product) CreateProduct(productEntry entry.ProductEntry) (*models.Products, error) {

	query := `INSERT INTO products (description, quantity) 
		VALUES (?, ?)`
	result, err := dao.DB.Exec(query, &productEntry.Description, &productEntry.Quantity)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return dao.GetProductByID(lastInsertId)
}
