package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/tkitsunai/edinet-go/api/edinet"
	v1 "github.com/tkitsunai/edinet-go/api/edinet/api/v1"
	"net/http"
)

type Documents struct {
	resourcePath string
	client       *edinet.V1Client
}

func NewDocumentsResource() *Documents {
	return &Documents{
		resourcePath: "documents",
		client:       edinet.NewV1Client(),
	}
}

func (d Documents) GetDocuments() (method, sPath string, fn func(ctx *gin.Context)) {
	return http.MethodGet, d.resourcePath, func(ctx *gin.Context) {
		p := FileDateParam{}
		ctx.ShouldBind(&p)

		res, err := d.client.RequestDocumentList(v1.FileDate(p.From))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		}

		ctx.JSON(http.StatusOK, res)
	}
}

type FileDateParam struct {
	From string `form:"date"`
}
