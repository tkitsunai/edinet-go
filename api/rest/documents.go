package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/tkitsunai/edinet-go/api/domain"
	"github.com/tkitsunai/edinet-go/api/edinet"
	v1 "github.com/tkitsunai/edinet-go/api/edinet/api/v1"
	"github.com/tkitsunai/edinet-go/api/usecase"
	"net/http"
)

const (
	DocumentResourcesRoot     = "documents"
	TermDocumentResourcesRoot = "termdocuments"
)

type Documents struct {
	resourcePath string
	client       *edinet.V1Client
	useCase      usecase.OverviewTerm
}

func NewDocumentsResource() *Documents {
	cli := edinet.NewV1Client()
	return &Documents{
		resourcePath: DocumentResourcesRoot,
		client:       cli,
		useCase: usecase.OverviewTerm{
			Client: cli,
		},
	}
}

func (d Documents) GetDocuments() (method, sPath string, fn func(ctx *gin.Context)) {
	return http.MethodGet, d.resourcePath, func(ctx *gin.Context) {
		p := FileDateParam{}
		err := ctx.ShouldBind(&p)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res, err := d.client.RequestDocumentList(v1.FileDate(p.From))

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (d Documents) GetDocumentsByTerm() (method, sPath string, fn func(ctx *gin.Context)) {
	return http.MethodGet, TermDocumentResourcesRoot, func(ctx *gin.Context) {
		p := FileTermParam{}
		err := ctx.ShouldBind(&p)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		from := v1.FileDate(p.From)
		to := v1.FileDate(p.To)
		_, err = from.Validate()
		_, err = to.Validate()

		term := domain.Term{
			FromDate: from,
			ToDate:   to,
		}

		overviews, errs := d.useCase.FindOverviewByTerm(term)

		if errs != nil && len(errs) > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
			ctx.Abort()
		}

		var res []v1.DocumentListResponse
		for _, overview := range overviews {
			if overview != nil {
				res = append(res, *overview)
			}
		}

		ctx.JSON(http.StatusOK, v1.DocumentListResponses{
			List: res,
		})
	}
}

type FileDateParam struct {
	From string `form:"date"`
}

type FileTermParam struct {
	From string `form:"from"`
	To   string `form:"to"`
}
