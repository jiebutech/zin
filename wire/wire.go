//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/jiebutech/datasource"
	"github.com/jiebutech/zin/app/api"
	"github.com/jiebutech/zin/app/repo"
	"github.com/jiebutech/zin/app/service"
)

func NewDemoController(orm *datasource.Orm) *api.DemoCtrl {
	panic(wire.Build(api.NewDemoCtrl, service.NewDemoSvc, repo.NewDemoRepo))
}
