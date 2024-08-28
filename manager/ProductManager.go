package manager

import (
	"GoPractice/dao"
	"GoPractice/entry"
	"GoPractice/models"
	"github.com/sirupsen/logrus"
)

type ProductManager struct {
	DAO *dao.Product
	Log *logrus.Logger
}

func NewProductManager(dao *dao.Product, log *logrus.Logger) *ProductManager {
	return &ProductManager{DAO: dao, Log: log}
}

func (m *ProductManager) GetProductByID(id int64) (*models.Products, error) {
	return m.DAO.GetProductByID(id)
}

func (m *ProductManager) CreateProduct(productEntry entry.ProductEntry) (*models.Products, error) {
	return m.DAO.CreateProduct(productEntry)
}
