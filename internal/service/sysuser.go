package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/liangtengfei/beer-admin/api/model"
	"github.com/liangtengfei/beer-admin/internal/biz"
	"go.uber.org/zap"
)

type SysUserService struct {
	UserCase *biz.SysUserUseCase
	Logger   *zap.SugaredLogger
}

func NewSysUserService(uc *biz.SysUserUseCase, logger *zap.SugaredLogger) *SysUserService {
	return &SysUserService{
		UserCase: uc,
		Logger:   logger.With(zap.String("module", "service/sysuser")),
	}
}

func (s *SysUserService) GetSysUser(ctx *gin.Context, req model.ByIdRequest) (model.SysUserInfo, error) {
	return s.UserCase.GetSysUser(ctx, req)
}

func (s *SysUserService) CreateSysUser(ctx *gin.Context, req model.CreateSysUserReq) (uint64, error) {
	existed := s.UserCase.ExistsSysUserByUserName(ctx, req.UserName, uint64(0))
	if existed {
		return uint64(0), errors.New("用户名已存在")
	}
	existed = s.UserCase.ExistsSysUserByMobile(ctx, req.Mobile, uint64(0))
	if existed {
		return uint64(0), errors.New("手机号码已存在")
	}

	data, err := s.UserCase.CreateSysUser(ctx, req)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *SysUserService) UpdateSysUser(ctx *gin.Context, req model.UpdateSysUserReq) (uint64, error) {
	existed := s.UserCase.ExistsSysUserByMobile(ctx, req.Mobile, req.ID)
	if existed {
		return uint64(0), errors.New("手机号码已存在")
	}

	data, err := s.UserCase.UpdateSysUser(ctx, req)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *SysUserService) DeleteSysUser(ctx *gin.Context, req model.ByIdRequest) (bool, error) {
	return s.UserCase.DeleteSysUser(ctx, req)
}

func (s *SysUserService) PageSysUser(ctx *gin.Context, req model.PageSysUserReq) ([]*model.SysUserInfo, int, error) {
	return s.UserCase.PageSysUser(ctx, req)
}
