package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liangtengfei/beer-admin/api/model"
	"github.com/liangtengfei/beer-admin/internal/conf"
	"github.com/liangtengfei/beer-admin/internal/service"
	"go.uber.org/zap"
)

var SysUserSet = wire.NewSet(wire.Struct(new(SysUserAPI), "*"))

type SysUserAPI struct {
	Conf    *conf.Server
	UserSrv *service.SysUserService
	Logger  *zap.SugaredLogger
}

// GetSysUser 用户详情
// @Summary      用户详情
// @Description  用户详情 by ID
// @Tags         用户: SysUser
// @Accept       json
// @Produce      json
// @Success      200  {object}  RestRes{data=model.SysUserInfo}
// @Router       /user/{id} [get]
func (a *SysUserAPI) GetSysUser(ctx *gin.Context) {
	var req model.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ErrorValid(ctx, err)
		return
	}

	data, err := a.UserSrv.GetSysUser(ctx, req)
	if err != nil {
		ErrorErr(ctx, err)
		return
	}

	SuccessData(ctx, data)
}

func (a *SysUserAPI) AddSysUser(ctx *gin.Context) {
	var req model.CreateSysUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorValid(ctx, err)
		return
	}

	data, err := a.UserSrv.CreateSysUser(ctx, req)
	if err != nil {
		ErrorErr(ctx, err)
		return
	}

	SuccessData(ctx, data)
}

func (a *SysUserAPI) EditSysUser(ctx *gin.Context) {
	var req model.UpdateSysUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorValid(ctx, err)
		return
	}

	data, err := a.UserSrv.UpdateSysUser(ctx, req)
	if err != nil {
		ErrorErr(ctx, err)
		return
	}

	SuccessData(ctx, data)
}

func (a *SysUserAPI) DelSysUser(ctx *gin.Context) {
	var req model.ByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ErrorValid(ctx, err)
		return
	}

	data, err := a.UserSrv.DeleteSysUser(ctx, req)
	if err != nil {
		ErrorErr(ctx, err)
		return
	}

	SuccessData(ctx, data)
}

func (a *SysUserAPI) PageSysUser(ctx *gin.Context) {
	var req model.PageSysUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorValid(ctx, err)
		return
	}
	
	data, total, err := a.UserSrv.PageSysUser(ctx, req)
	if err != nil {
		ErrorErr(ctx, err)
		return
	}

	SuccessPage(ctx, total, data)
}
