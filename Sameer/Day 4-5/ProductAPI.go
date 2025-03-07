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
func getProducts(c *gin.Context) {
	rows, err := db.Query("SELECT product_id, product_name, price, quantity FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []map[string]interface{}
	for rows.Next() {
		var product_id, price, quantity int
		var product_name string
		rows.Scan(&product_id, &product_name, &quantity, &price)
		books = append(books, gin.H{"product_id": product_id, "product_name": product_name, "price": price, "quantity": quantity})
	}

	c.JSON(http.StatusOK, books)
}
func updateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var update struct {
		Price    *int `json:"price"`
		Quantity *int `json:"quantity"`
	}
	if c.BindJSON(&update) != nil || (update.Price == nil || update.Quantity == nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	db.Exec("UPDATE products SET price = ?, quantity = ? WHERE product_id = ?", update.Price, update.Quantity, id)
	c.JSON(http.StatusOK, gin.H{"message": "Updated"})
}
func placeOrder(c *gin.Context) {
	var order struct {
		CustomerID int `json:"customer_id"`
		ProductID  int `json:"product_id"`
		Quantity   int `json:"quantity"`
	}

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AAH! , Invalid input"})
		return
	}

	var availableQuantity int
	var status string
	err := db.QueryRow("SELECT quantity FROM products WHERE product_id = ?", order.ProductID).Scan(&availableQuantity)
	if err != nil {
		status = "Failed"
		c.JSON(http.StatusNotFound, gin.H{"error": "Let me check, Oh product not found!"})
		saveOrder(db, order.CustomerID, order.ProductID, order.Quantity, status)
		return
	}

	if order.Quantity > availableQuantity {
		status = "Failed"
		c.JSON(http.StatusBadRequest, gin.H{"error": "Let me check, Oh not enough stocks!"})
		saveOrder(db, order.CustomerID, order.ProductID, order.Quantity, status)

		return
	}

	_, err = db.Exec("UPDATE products SET quantity = quantity - ? WHERE product_id = ?", order.Quantity, order.ProductID)
	if err != nil {
		status = "Failed"
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		saveOrder(db, order.CustomerID, order.ProductID, order.Quantity, status)

		return
	}
	status = "Order Placed"
	saveOrder(db, order.CustomerID, order.ProductID, order.Quantity, status)

	c.JSON(http.StatusCreated, gin.H{"message": "Yay Razors , Order placed successfully"})
}
func saveOrder(db *sql.DB, customerID, productID, quantity int, status string) {
	_, err := db.Exec("INSERT INTO orders (customer_id, product_id, quantity, order_status) VALUES (?, ?, ?, ?)", customerID, productID, quantity, status)
	if err != nil {
		fmt.Println("Failed to save order status:", err)
	}
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
	r.GET("/products", getProducts)
	r.PATCH("/product/:id", updateProduct)
	r.POST("/order", placeOrder)
	r.Run(":8080")
}
