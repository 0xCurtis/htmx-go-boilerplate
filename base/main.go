package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	// Initialize the database
	InitDB()

	// Define HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/listUser", listUserHandler)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("template/listUser.html"))
	tmpl.Execute(w, users)
}
