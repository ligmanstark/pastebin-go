package service

import (
	"database/sql"
	v2 "ligmanstark/pastebin_v2/model/v2"
	pkg "ligmanstark/pastebin_v2/packages"
	"os"
	"path/filepath"
	"time"
)

type CreateTextPastebinRequest struct {
	Content string `json:"content" binding:"required"`
}

type CreateImagePastebinRequest struct {
	ImageData []byte `json:"image_data" binding:"required"`
	MimeType  string `json:"mime_type"`
}

type CreateResponse struct {
	UrlSlug string `json:"url_slug"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TextPastebin interface {
	CreateTextPastebin(content string) (CreateResponse, ErrorResponse)
	GetTextPastebinBySlug(slug string) (v2.TextPastebin, ErrorResponse)
	GetAllTextPastebins() ([]v2.TextPastebin, ErrorResponse)
}

type ImagePastebin interface {
	CreateImagePastebin(image_data []byte, fileSize int, mimeType string) (CreateResponse, ErrorResponse)
	GetImagePastebinBySlug(slug string) (v2.ImagePastebin, ErrorResponse)
	GetAllImagePastebins() ([]v2.ImagePastebin, ErrorResponse)
}

const UPLOAD_DIR = "./uploads"

type PastebinService struct {
	db *sql.DB
}

func NewPastebinService(db *sql.DB) *PastebinService {
	return &PastebinService{
		db: db,
	}
}

func (service *PastebinService) CreateTextPastebin(content string) (CreateResponse, ErrorResponse) {
	if content == "" {
		return CreateResponse{}, ErrorResponse{Code: 400, Message: "Параметр content не может быть пустым"}
	}

	slug, err := pkg.GenerateRandomSlug(12)
	if err != nil {
		return CreateResponse{}, ErrorResponse{Code: 500, Message: "Ошибка генерации URL: " + err.Error()}
	}
	pastebin := v2.TextPastebin{
		UrlSlug:   slug,
		Content:   content,
		CreatedAt: time.Now(),
	}
	query := "INSERT INTO text_pastebin (url_slug, content, created_at) VALUES ($1, $2, $3)"
	_, err = service.db.Exec(query, pastebin.UrlSlug, pastebin.Content, pastebin.CreatedAt)
	if err != nil {
		return CreateResponse{}, ErrorResponse{Code: 500, Message: "Ошибка сохранения в базе данных: " + err.Error()}
	}
	return CreateResponse{UrlSlug: pastebin.UrlSlug}, ErrorResponse{}
}

func (service *PastebinService) GetTextPastebinBySlug(slug string) (v2.TextPastebin, ErrorResponse) {
	var pastebin v2.TextPastebin
	query := "SELECT id, url_slug, created_at, content FROM text_pastebin WHERE url_slug = $1"
	err := service.db.QueryRow(query, slug).Scan(&pastebin.ID, &pastebin.UrlSlug, &pastebin.CreatedAt, &pastebin.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return v2.TextPastebin{}, ErrorResponse{Code: 404, Message: "Пост не найден"}
		}
		return v2.TextPastebin{}, ErrorResponse{Code: 500, Message: "Ошибка запроса к базе данных: " + err.Error()}
	}
	return pastebin, ErrorResponse{}
}

func (service *PastebinService) GetAllTextPastebins() ([]v2.TextPastebin, ErrorResponse) {
	var pastebins []v2.TextPastebin
	query := "SELECT id, url_slug, created_at, content FROM text_pastebin"
	rows, err := service.db.Query(query)
	if err != nil {
		return nil, ErrorResponse{Code: 500, Message: "Ошибка запроса к базе данных: " + err.Error()}
	}
	defer rows.Close()

	for rows.Next() {
		var pastebin v2.TextPastebin
		err := rows.Scan(&pastebin.ID, &pastebin.UrlSlug, &pastebin.CreatedAt, &pastebin.Content)
		if err != nil {
			return nil, ErrorResponse{Code: 500, Message: "Ошибка чтения данных из базы данных: " + err.Error()}
		}
		pastebins = append(pastebins, pastebin)
	}
	return pastebins, ErrorResponse{}
}

func (service *PastebinService) CreateImagePastebin(image_data []byte, fileSize int, mimeType string) (CreateResponse, ErrorResponse) {
	if len(image_data) == 0 {
		return CreateResponse{}, ErrorResponse{Code: 400, Message: "Изображение не может быть пустым"}
	}

	slug, err := pkg.GenerateRandomSlug(12)
	if err != nil {
		return CreateResponse{}, ErrorResponse{Code: 500, Message: "Ошибка генерации URL: " + err.Error()}
	}

	err = os.MkdirAll(UPLOAD_DIR, os.ModePerm)
	if err != nil {
		return CreateResponse{}, ErrorResponse{Code: 500, Message: "Ошибка создания директории для загрузки: " + err.Error()}
	}

	ext := pkg.GetExtensionMime(mimeType)
	filePath := filepath.Join(UPLOAD_DIR, slug+ext)

	err = os.WriteFile(filePath, image_data, 0644)
	if err != nil {
		return CreateResponse{}, ErrorResponse{Code: 500, Message: "Ошибка сохранения файла: " + err.Error()}
	}

	query := "INSERT INTO Image_Pastebin (url_slug,file_size, mime_type, created_at) VALUES ($1, $2, $3, $4)"
	_, err = service.db.Exec(query, slug, fileSize, mimeType, time.Now())

	if err != nil {
		return CreateResponse{}, ErrorResponse{Code: 500, Message: "Ошибка сохранения в базе данных: " + err.Error()}
	}

	return CreateResponse{UrlSlug: slug}, ErrorResponse{}
}

func (service *PastebinService) GetImagePastebinBySlug(slug string) (v2.ImagePastebin, ErrorResponse) {
	var pastebin v2.ImagePastebin
	query := "SELECT id, url_slug, file_size, mime_type, created_at FROM Image_Pastebin WHERE url_slug = $1"
	err := service.db.QueryRow(query, slug).Scan(&pastebin.ID, &pastebin.UrlSlug, &pastebin.FileSize, &pastebin.MimeType, &pastebin.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return v2.ImagePastebin{}, ErrorResponse{Code: 404, Message: "Изображение не найдено"}
		}
		return v2.ImagePastebin{}, ErrorResponse{Code: 500, Message: "Ошибка запроса к базе данных: " + err.Error()}
	}

	ext := pkg.GetExtensionMime(pastebin.MimeType)
	filePath := filepath.Join(UPLOAD_DIR, slug+ext)
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return v2.ImagePastebin{}, ErrorResponse{Code: 404, Message: "Изображение не найдено: " + err.Error()}
	}
	pastebin.ImageData = imageData
	return pastebin, ErrorResponse{}
}
