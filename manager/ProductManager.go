package manager

import (
	"GoPractice/database"
	"GoPractice/entry"
	"GoPractice/models"
	"github.com/sirupsen/logrus"
)

type ProductManager struct {
	DAO *database.Product
	Log *logrus.Logger
}

func NewProductManager(dao *database.Product, log *logrus.Logger) *ProductManager {
	return &ProductManager{DAO: dao, Log: log}
}

func (m *ProductManager) GetProductByID(id int64) (*models.Products, error) {
	return m.DAO.GetProductByID(id)
}

func (m *ProductManager) CreateProduct(productEntry entry.ProductEntry) (*models.Products, error) {
	return m.DAO.CreateProduct(productEntry)
}
