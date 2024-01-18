package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebutech/zin/app/api"
)

// RegisterDemoRoute 路app
func RegisterDemoRoute(eng *gin.Engine, ctrl *api.DemoCtrl) {
	eng.GET("/demo", ctrl.GetDemoById)
}
