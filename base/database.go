package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// User struct represents a user in the system
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}

// InitDB initializes the database connection and migrates the schema
func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	// Seed initial data
	seedUsers()
}

// seedUsers inserts initial user data into the database
func seedUsers() {
	insertUser("John Doe", "john@example.com")
	insertUser("Jane Smith", "jane@example.com")
}

// insertUser inserts a new user into the database
func insertUser(name, email string) {
	user := User{Name: name, Email: email}
	db.FirstOrCreate(&user, User{Email: email})
}
