package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebutech/app"
	"github.com/jiebutech/zin/app/dto"
	"github.com/jiebutech/zin/app/service"
)

type DemoCtrl struct {
	svc *service.DemoSvc
}

func NewDemoCtrl(svc *service.DemoSvc) *DemoCtrl {
	return &DemoCtrl{svc: svc}
}

func (n DemoCtrl) GetDemoById(c *gin.Context) {
	ctx := app.ConvertToApiContext(c)
	req := new(dto.IdReq)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.ResponseErrorCode(app.ErrBadRequest, err.Error())
		return
	}
	info, err := n.svc.GetById(ctx, req.Id)
	if err != nil {
		ctx.ResponseError(err)
		return
	}
	// convert info to the response struct
	ctx.SuccessData(info)
}
