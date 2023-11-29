package di

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/conf"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/gateway"
	"github.com/tkitsunai/edinet-go/usecase"
)

func GetDIContainer() *do.Injector {
	injector := do.New()

	config := conf.LoadConfig()

	do.ProvideValue(injector, config)
	do.Provide(injector, edinet.NewClient)

	do.Provide(injector, gateway.NewDocument)
	do.Provide(injector, gateway.NewOverview)

	do.Provide(injector, usecase.NewOverview)
	do.Provide(injector, usecase.NewDocument)

	return injector
}
