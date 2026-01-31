package handlers

import (
	"encoding/json"
	"ligmanstark/pastebin-go/model"
	"ligmanstark/pastebin-go/services"
	"net/http"
)

func CreatePastebinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var pastebin model.Pastebin
	if err := json.NewDecoder(r.Body).Decode(&pastebin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := services.InitDB()
	defer db.Close()

	generatedSlug := services.GenerateRandomSlug()
	pastebin.UrlSlug = generatedSlug

	_, err := db.Exec("INSERT INTO pastebin (content, url_slug) VALUES ($1, $2)", pastebin.Content, pastebin.UrlSlug)
	if err != nil {
		http.Error(w, "Failed creating pastebin: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pastebin)
}

func GetPastebinBySlug(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	urlSlug := r.URL.Query().Get("id")
	db := services.InitDB()
	defer db.Close()

	var pastebin model.Pastebin
	err := db.QueryRow("SELECT id, content, created_at, url_slug FROM pastebin WHERE url_slug = $1", urlSlug).Scan(&pastebin.ID, &pastebin.Content, &pastebin.CreatedAt, &pastebin.UrlSlug)
	if err != nil {
		http.Error(w, "Failed to retrieve pastebin: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pastebin)
}
