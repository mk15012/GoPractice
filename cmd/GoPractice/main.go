package main

import (
	"GoPractice/config"
	"GoPractice/controller"
	"GoPractice/database"
	"GoPractice/manager"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	// Initialize logrus
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Load configuration
	configPath := filepath.Join("profiles", "stage", "conf", "config-prod.yml")
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Build the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Datastore.MySQL.UserName, cfg.Datastore.MySQL.Password, cfg.Datastore.MySQL.Host, cfg.Datastore.MySQL.Port, cfg.Datastore.MySQL.DbName)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	routes := mux.NewRouter()

	userDAO := database.NewUser(db)
	userManager := manager.NewUserManager(userDAO, logger)
	userController := controller.NewUserController(userManager, logger)

	productDAO := database.NewProduct(db)
	productManager := manager.NewProductManager(productDAO, logger)
	productController := controller.NewProductController(productManager, logger)

	// Set up HTTP handlers for users
	routes.HandleFunc("/users/getUser/{userId}", userController.GetUserById).Methods("GET")
	routes.HandleFunc("/users/createUser", userController.CreateUser).Methods("POST")
	routes.HandleFunc("/users/updateUser", userController.UpdateUser).Methods("PUT")
	routes.HandleFunc("/users/deleteUser/{userId}", userController.DeleteUser).Methods("DELETE")

	// Set up HTTP handlers for products
	routes.HandleFunc("/products/getProduct/{productId}", productController.GetProductById).Methods("GET")
	routes.HandleFunc("/products/createProduct", productController.CreateProduct).Methods("POST")
	routes.HandleFunc("/products/updateProduct", productController.UpdateProduct).Methods("PUT")
	routes.HandleFunc("/products/deleteProduct/{productId}", productController.DeleteProduct).Methods("DELETE")

	// Start the server
	httpPort := fmt.Sprintf(":%d", cfg.Port.HTTPPort)
	log.Fatal(http.ListenAndServe(httpPort, routes))

}
