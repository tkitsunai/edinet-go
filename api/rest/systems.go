package rest

import (
	"github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Systems struct {
	resourcePath string
}

func NewSystemsResource() *Systems {
	return &Systems{
		resourcePath: "",
	}
}

func (d Systems) Systems() (method, sPath string, fn func(ctx *gin.Context)) {
	return http.MethodGet, d.resourcePath, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"systems": stats_api.GetStats(),
		})
	}
}
