package v2

import (
	"fmt"
	"io"
	pkg "ligmanstark/pastebin_v2/packages"
	"ligmanstark/pastebin_v2/service"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var pastebinService *service.PastebinService

func SetPastebinService(svc *service.PastebinService) {
	pastebinService = svc
}
func CreateTextPastebinHandler(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: "Ошибка чтения тела запроса: " + err.Error()})
		return
	}
	content := string(body)
	if content == "" {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: "Параметр content не может быть пустым"})
		return
	}

	response, errResp := pastebinService.CreateTextPastebin(content)
	if errResp.Code != 0 {
		ctx.JSON(errResp.Code, errResp)
		return
	}
	ctx.JSON(301, response)
}

func GetTextPastebinBySlugHandler(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, errResp := pastebinService.GetTextPastebinBySlug(slug)
	if errResp.Code != 0 {
		ctx.JSON(errResp.Code, errResp)
		return
	}
	ctx.JSON(200, response)
}

func GetTextPastebinAllHandler(ctx *gin.Context) {
	response, errResp := pastebinService.GetAllTextPastebins()
	if errResp.Code != 0 {
		ctx.JSON(errResp.Code, errResp)
		return
	}
	ctx.JSON(200, response)
}

func CreateImagePastebinHandler(ctx *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		size = os.Getenv("MAX_SIZE_FILE")
	)
	maxSizeFile, err := strconv.Atoi(size)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: "Файл не загружен: " + err.Error()})
		return
	}
	if file.Size > int64(maxSizeFile)*1024*1024 {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: fmt.Sprintf("Размер файла превышает %d МБ", maxSizeFile)})
		return
	}
	if file.Size == 0 {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: "Файл не может быть пустым"})
		return
	}

	fileData, err := file.Open()
	if err != nil {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: "Ошибка открытия файла: " + err.Error()})
		return
	}
	defer fileData.Close()

	imageBytes := make([]byte, file.Size)
	_, err = fileData.Read(imageBytes)
	if err != nil {
		ctx.JSON(400, service.ErrorResponse{Code: 400, Message: "Ошибка чтения файла: " + err.Error()})
		return
	}

	mimeType := file.Header.Get("Content-Type")
	response, errResp := pastebinService.CreateImagePastebin(imageBytes, int(file.Size), mimeType)
	if errResp.Code != 0 {
		ctx.JSON(errResp.Code, errResp)
		return
	}
	ctx.JSON(301, response)
}

func GetImagePastebinBySlugHandler(ctx *gin.Context) {
	slug := ctx.Param("slug")
	response, errResp := pastebinService.GetImagePastebinBySlug(slug)
	if errResp.Code != 0 {
		ctx.JSON(errResp.Code, errResp)
		return
	}

	ctx.Header("Content-Type", response.MimeType)
	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", slug+pkg.GetExtensionMime(response.MimeType)))
	ctx.Data(200, response.MimeType, response.ImageData)
}
