package di

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/conf"
	"github.com/tkitsunai/edinet-go/datastore"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/gateway"
	"github.com/tkitsunai/edinet-go/server"
	"github.com/tkitsunai/edinet-go/usecase"
)

func SetUpContainer(storeDriver datastore.Driver) *do.Injector {
	injector := do.New()

	config := conf.LoadConfig()

	do.ProvideValue(injector, config)
	do.ProvideValue[datastore.Driver](injector, storeDriver)
	do.Provide(injector, edinet.NewClient)

	do.Provide(injector, gateway.NewDocument)
	do.Provide(injector, gateway.NewCompany)
	do.Provide(injector, gateway.NewOverview)

	do.Provide(injector, usecase.NewOverview)
	do.Provide(injector, usecase.NewDocument)
	do.Provide(injector, usecase.NewCompany)

	do.Provide(injector, server.NewDocuments)
	do.Provide(injector, server.NewCompany)

	return injector
}
