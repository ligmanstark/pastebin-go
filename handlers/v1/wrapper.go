package v1

import (
	"github.com/gin-gonic/gin"
)

func CreatePastebinHandlerWrapper(ctx *gin.Context) {
	req := ctx.Request
	CreatePastebinHandler(ctx.Writer, req)
}

func GetPastebinBySlugHandlerWrapper(ctx *gin.Context) {
	slug := ctx.Param("slug")
	req := ctx.Request
	if slug != "" && req != nil && req.URL != nil {
		query := req.URL.Query()
		query.Set("id", slug)
	}
	GetPastebinBySlugHandler(ctx.Writer, req)
}

func GetPastebinAllHandlerWrapper(ctx *gin.Context) {
	req := ctx.Request
	GetPastebinAllHandler(ctx.Writer, req)
}
