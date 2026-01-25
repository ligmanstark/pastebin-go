package handlers

import (
	"encoding/json"
	"fmt"
	"ligmanstark/pastebin-go/model"
	"ligmanstark/pastebin-go/services"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db := services.InitDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users (name, role_id) VALUES ($1, $2)", user.Name, user.RoleID)
	if err != nil {
		http.Error(w, "Failed creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User created successfully")
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db := services.InitDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Failed to retrieve users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.RoleID); err != nil {
			http.Error(w, "Failed to scan user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
