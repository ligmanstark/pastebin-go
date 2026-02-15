package v1

import (
	"encoding/json"
	"io"
	model_v1 "ligmanstark/pastebin_v2/model/v1"
	services "ligmanstark/pastebin_v2/packages"
	"net/http"
)

func CreatePastebinHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var pastebin model_v1.Pastebin
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pastebin.Content = string(body)
	db := services.InitDB()
	defer db.Close()

	generatedSlug, err := services.GenerateRandomSlug(8)
	if err != nil {
		http.Error(w, "Failed to generate URL slug: "+err.Error(), http.StatusInternalServerError)
		return
	}
	pastebin.UrlSlug = generatedSlug

	_, err = db.Exec("INSERT INTO pastebin (content, url_slug) VALUES ($1, $2)", pastebin.Content, pastebin.UrlSlug)
	if err != nil {
		http.Error(w, "Failed creating pastebin: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pastebin)
}

func GetPastebinBySlugHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	urlSlug := r.URL.Query().Get("id")
	db := services.InitDB()
	defer db.Close()

	var pastebin model_v1.Pastebin
	err := db.QueryRow("SELECT id, content, created_at, url_slug FROM pastebin WHERE url_slug = $1", urlSlug).Scan(&pastebin.ID, &pastebin.Content, &pastebin.CreatedAt, &pastebin.UrlSlug)
	if err != nil {
		http.Error(w, "Failed to retrieve pastebin: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pastebin)
}

func GetPastebinAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	db := services.InitDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pastebin")
	if err != nil {
		http.Error(w, "Failed to retrieve pastebins: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pastebins []model_v1.Pastebin
	for rows.Next() {
		var pastebin model_v1.Pastebin
		if err := rows.Scan(&pastebin.ID, &pastebin.Content, &pastebin.CreatedAt, &pastebin.UrlSlug); err != nil {
			http.Error(w, "Failed to scan pastebin: "+err.Error(), http.StatusInternalServerError)
			return
		}
		pastebins = append(pastebins, pastebin)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pastebins)
}
