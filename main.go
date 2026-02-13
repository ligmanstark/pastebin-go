package main

import (
	"log"

	pkg "ligmanstark/pastebin_v2/packages"

	"os"

	v1 "ligmanstark/pastebin_v2/handlers/v1"
	v2 "ligmanstark/pastebin_v2/handlers/v2"
	"ligmanstark/pastebin_v2/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := pkg.InitDB()
	pastebinService := service.NewPastebinService(db)
	v2.SetPastebinService(pastebinService)
}

func main() {

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	router := gin.Default()
	router.Use(cors.Default())

	v1Group := router.Group("/api/v1")
	{
		v1Group.POST("/pastebin/create", v1.CreatePastebinHandlerWrapper)
		v1Group.GET("/pastebin/:slug", v1.GetPastebinBySlugHandlerWrapper)
		v1Group.GET("/pastebin/all", v1.GetPastebinAllHandlerWrapper)
	}

	v2Group := router.Group("/api/v2")
	{
		v2Group.POST("/text/create", v2.CreateTextPastebinHandler)
		v2Group.GET("/text/:slug", v2.GetTextPastebinBySlugHandler)
		v2Group.GET("/text/all", v2.GetTextPastebinAllHandler)

		v2Group.POST("/image/create", v2.CreateImagePastebinHandler)
		v2Group.GET("/image/:slug", v2.GetImagePastebinBySlugHandler)

	}

	log.Fatal(router.Run(appHost + ":" + appPort))

}
