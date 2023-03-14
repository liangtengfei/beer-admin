package biz

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/liangtengfei/beer-admin/api/model"
	"github.com/liangtengfei/beer-admin/internal/conf"
	"github.com/liangtengfei/beer-admin/internal/data/ent"
	"github.com/liangtengfei/beer-admin/pkg/jwt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type SysUserRepo interface {
	GetSysUser(context.Context, model.ByIdRequest) (*ent.SysUser, error)
	GetSysUserByUserNameAll(context.Context, string) (*ent.SysUser, error)
	GetSysUserByMobileAll(context.Context, string) (*ent.SysUser, error)
	CreateSysUser(context.Context, model.CreateSysUserReq, string) (*ent.SysUser, error)
	UpdateSysUser(context.Context, model.UpdateSysUserReq, string) (*ent.SysUser, error)
	DeleteSysUser(context.Context, model.ByIdRequest, string) (bool, error)
	PageSysUser(context.Context, model.PageSysUserReq) ([]*ent.SysUser, int, error)
}
type SysUserUseCase struct {
	repo  SysUserRepo
	conf  *conf.Server
	cache redis.Cmdable
	log   *zap.SugaredLogger
}

func NewSysUserUseCase(c *conf.Server, r SysUserRepo, cache redis.Cmdable, logger *zap.SugaredLogger) *SysUserUseCase {
	return &SysUserUseCase{
		repo:  r,
		conf:  c,
		cache: cache,
		log:   logger.With(zap.String("module", "usecase/sysuser")),
	}
}

func (uc *SysUserUseCase) GetSysUser(ctx *gin.Context, req model.ByIdRequest) (model.SysUserInfo, error) {
	data, err := uc.repo.GetSysUser(ctx, req)
	if err != nil {
		uc.log.Errorf("获取用户信息失败：%v", err)
		return model.SysUserInfo{}, errors.New("未查询到数据")
	}

	postIds := make([]uint64, 0)
	for _, post := range data.Edges.UsersPost {
		postIds = append(postIds, post.ID)
	}

	var reply model.SysUserInfo
	err = structCopy(uc.log, &reply, &data)
	if err != nil {
		return reply, err
	}

	reply.PostIds = postIds

	return reply, nil
}

// ExistsSysUserByUserName 检查用户名是否存在
func (uc *SysUserUseCase) ExistsSysUserByUserName(ctx *gin.Context, req string, id uint64) bool {
	data, err := uc.repo.GetSysUserByUserNameAll(ctx, req)
	// 如果查找数据库失败 返回true 禁止数据插入
	if err != nil {
		if ent.IsNotFound(err) {
			return false
		}
		uc.log.Errorf("检查用户名是否唯一发生错误：%v", err)
		return true
	}

	if data != nil && data.ID > 0 && id > 0 && data.ID != id {
		return true
	}

	return false
}

func (uc *SysUserUseCase) ExistsSysUserByMobile(ctx *gin.Context, req string, id uint64) bool {
	data, err := uc.repo.GetSysUserByMobileAll(ctx, req)
	if err != nil {
		if ent.IsNotFound(err) {
			return false
		}
		uc.log.Errorf("检查手机号码是否唯一发生错误：%v", err.Error())
		return true
	}

	if data != nil && data.ID > 0 && id > 0 && data.ID != id {
		return true
	}

	return false
}

func (uc *SysUserUseCase) CreateSysUser(ctx *gin.Context, req model.CreateSysUserReq) (uint64, error) {
	username, err := jwt.LoginUserName(ctx)
	if err != nil {
		return uint64(0), err
	}

	data, err := uc.repo.CreateSysUser(ctx, req, username)
	if err != nil {
		uc.log.Errorf("新增用户失败：%v", err)
		return uint64(0), errors.New("新增用户失败")
	}

	return data.ID, nil
}

func (uc *SysUserUseCase) UpdateSysUser(ctx *gin.Context, req model.UpdateSysUserReq) (uint64, error) {
	username, err := jwt.LoginUserName(ctx)
	if err != nil {
		return uint64(0), err
	}

	data, err := uc.repo.UpdateSysUser(ctx, req, username)
	if err != nil {
		uc.log.Errorf("更新用户信息失败：%v", err)
		return uint64(0), errors.New("更新用户信息失败")
	}

	var reply model.SysUserInfo
	err = structCopy(uc.log, &reply, &data)
	if err != nil {
		uc.log.Errorf("更新用户信息失败：%v", err)
		return uint64(0), errors.New("更新用户信息失败")
	}

	return data.ID, nil
}

func (uc *SysUserUseCase) DeleteSysUser(ctx *gin.Context, req model.ByIdRequest) (bool, error) {
	username, err := jwt.LoginUserName(ctx)
	if err != nil {
		return false, err
	}

	_, err = uc.repo.DeleteSysUser(ctx, req, username)
	if err != nil {
		uc.log.Errorf("删除用户信息失败：%v", err)
		return false, errors.New("删除用户信息失败")
	}
	return true, nil
}

func (uc *SysUserUseCase) PageSysUser(ctx *gin.Context, req model.PageSysUserReq) ([]*model.SysUserInfo, int, error) {
	data, total, err := uc.repo.PageSysUser(ctx, req)
	// 把错误信息转为可以对外展示的信息
	if err != nil {
		uc.log.Errorf("分页查询用户信息失败：%v", err)
		return nil, 0, errors.New("分页查询用户信息失败")
	}

	var reply []*model.SysUserInfo
	err = structCopy(uc.log, &reply, &data)
	if err != nil {
		return nil, 0, err
	}

	return reply, total, nil
}
