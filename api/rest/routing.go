package rest

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	g := gin.Default()

	g.Handle(NewSystemsResource().Systems())

	v1 := g.Group("v1")

	doc := NewDocumentsResource()
	v1.Handle(doc.GetDocuments())
	v1.Handle(doc.GetDocumentsByTerm())

	router = g
}

func Run() error {
	return router.Run()
}
