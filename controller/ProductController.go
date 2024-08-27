package controller

import (
	"GoPractice/entry"
	"GoPractice/manager"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ProductController struct {
	Manager *manager.ProductManager
	Log     *logrus.Logger
}

func NewProductController(manager *manager.ProductManager, log *logrus.Logger) *ProductController {
	return &ProductController{Manager: manager, Log: log}
}

func (c *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId, err := strconv.ParseInt(vars["productId"], 10, 64)

	if err != nil {
		c.Log.Errorf("Error parsing input : %v", err)
		http.Error(w, "Invalid productId", http.StatusBadRequest)
		return
	}

	product, err := c.Manager.GetProductByID(productId)
	if err != nil {
		c.Log.Errorf("Error fetching data : %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productEntry entry.ProductEntry
	err := json.NewDecoder(r.Body).Decode(&productEntry)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	product, err := c.Manager.CreateProduct(productEntry)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating product
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting product
}
