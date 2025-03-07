package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

func connectDatabase() {
	var err error
	dsn := "root:123456789@tcp(127.0.0.1:3306)/booksdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to MySQL!")
}

func addProduct(c *gin.Context) {
	var product struct {
		ID       int    `json:"product_id"`
		Name     string `json:"product_name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
	}

	if err := c.BindJSON(&product); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("INSERT INTO products (product_id,product_name, price, quantity) VALUES (?,?, ?, ?)", product.ID, product.Name, product.Price, product.Quantity)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product_id": product.ID, "product_name": product.Name, "price": product.Price, "quantity": product.Quantity, "message": "Product added successfully"})
}
func main() {
	connectDatabase()

	r := gin.Default()

	r.POST("/products", addProduct)
	r.Run(":8080")
}
