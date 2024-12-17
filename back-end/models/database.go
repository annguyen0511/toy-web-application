package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DB struct {
	*sql.DB
}

var GlobalDB *DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	GlobalDB = &DB{sqlDB}
	if err = GlobalDB.Ping(); err != nil {
		log.Fatal(err)
	}

	InitDB()
}

func InitDB() {
	// Tạo bảng Users
	_, err := GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Users (
		UserID INT PRIMARY KEY AUTO_INCREMENT,
		FullName VARCHAR(100) NOT NULL,
		Email VARCHAR(100) UNIQUE NOT NULL,
		PasswordHash VARCHAR(255) NOT NULL,
		PhoneNumber VARCHAR(15),
		Address TEXT,
		Role VARCHAR(50) NOT NULL DEFAULT 'User',
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatal("Error creating Users table: ", err)
	}

	// Tạo bảng Categories
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Categories (
		CategoryID INT PRIMARY KEY AUTO_INCREMENT,
		CategoryName VARCHAR(100) NOT NULL,
		Description TEXT,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatal("Error creating Categories table: ", err)
	}

	// Tạo bảng Products
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Products (
		ProductID INT PRIMARY KEY AUTO_INCREMENT,
		ProductName VARCHAR(100) NOT NULL,
		Description TEXT,
		Price DECIMAL(10, 2) NOT NULL,
		Stock INT NOT NULL,
		CategoryID INT,
		ImageURL VARCHAR(255),
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
	);`)
	if err != nil {
		log.Fatal("Error creating Products table: ", err)
	}

	// Tạo bảng Cart
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Cart (
		CartID INT PRIMARY KEY AUTO_INCREMENT,
		UserID INT NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (UserID) REFERENCES Users(UserID)
	);`)
	if err != nil {
		log.Fatal("Error creating Cart table: ", err)
	}

	// Tạo bảng CartDetails
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS CartDetails (
		CartDetailID INT PRIMARY KEY AUTO_INCREMENT,
		CartID INT NOT NULL,
		ProductID INT NOT NULL,
		Quantity INT NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (CartID) REFERENCES Cart(CartID),
		FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
	);`)
	if err != nil {
		log.Fatal("Error creating CartDetails table: ", err)
	}

	// Tạo bảng Orders
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Orders (
		OrderID INT PRIMARY KEY AUTO_INCREMENT,
		UserID INT NOT NULL,
		OrderDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		TotalAmount DECIMAL(10, 2) NOT NULL,
		OrderStatus VARCHAR(50) DEFAULT 'Pending',
		ShippingAddress TEXT,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (UserID) REFERENCES Users(UserID)
	);`)
	if err != nil {
		log.Fatal("Error creating Orders table: ", err)
	}

	// Tạo bảng OrderDetails
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS OrderDetails (
		OrderDetailID INT PRIMARY KEY AUTO_INCREMENT,
		OrderID INT NOT NULL,
		ProductID INT NOT NULL,
		Quantity INT NOT NULL,
		Price DECIMAL(10, 2) NOT NULL,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
		FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
	);`)
	if err != nil {
		log.Fatal("Error creating OrderDetails table: ", err)
	}

	// Tạo bảng Payments
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Payments (
		PaymentID INT PRIMARY KEY AUTO_INCREMENT,
		OrderID INT NOT NULL,
		mentDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		Amount DECIMAL(10, 2) NOT NULL,
		PaymentMethod VARCHAR(50) NOT NULL,
		PaymentStatus VARCHAR(50) DEFAULT 'Completed',
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (OrderID) REFERENCES Orders(OrderID)
	);`)
	if err != nil {
		log.Fatal("Error creating Payments table: ", err)
	}

	// Tạo bảng Reviews
	_, err = GlobalDB.Exec(`
	CREATE TABLE IF NOT EXISTS Reviews (
		ReviewID INT PRIMARY KEY AUTO_INCREMENT,
		ProductID INT NOT NULL,
		UserID INT NOT NULL,
		Rating INT CHECK (Rating BETWEEN 1 AND 5),
		ReviewText TEXT,
		ReviewDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
		FOREIGN KEY (UserID) REFERENCES Users(UserID)
	);`)
	if err != nil {
		log.Fatal("Error creating Reviews table: ", err)
	}
}
