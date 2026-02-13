package v1

import (
	"github.com/gin-gonic/gin"
)

func CreatePastebinHandlerWrapper(ctx *gin.Context) {
	req := ctx.Request
	CreatePastebinHandler(ctx.Writer, req)
}

func GetPastebinBySlugHandlerWrapper(ctx *gin.Context) {
	req := ctx.Request
	GetPastebinBySlugHandler(ctx.Writer, req)
}

func GetPastebinAllHandlerWrapper(ctx *gin.Context) {
	req := ctx.Request
	GetPastebinAllHandler(ctx.Writer, req)
}
