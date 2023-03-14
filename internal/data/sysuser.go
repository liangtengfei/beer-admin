package data

import (
	"context"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/liangtengfei/beer-admin/api/model"
	"github.com/liangtengfei/beer-admin/internal/biz"
	"github.com/liangtengfei/beer-admin/internal/data/ent"
	"github.com/liangtengfei/beer-admin/internal/data/ent/sysuser"
	"github.com/liangtengfei/beer-admin/pkg/password"
	"go.uber.org/zap"
)

var _ biz.SysUserRepo = (*sysUserRepo)(nil)

type sysUserRepo struct {
	data *Data
	log  *zap.SugaredLogger
}

func NewSysUserRepo(data *Data, logger *zap.SugaredLogger) biz.SysUserRepo {
	return &sysUserRepo{
		data: data,
		log:  logger.With(zap.String("module", "data/sysuser")),
	}
}

func (r *sysUserRepo) GetSysUser(ctx context.Context, req model.ByIdRequest) (*ent.SysUser, error) {
	return r.data.EntDb.SysUser.
		Query().
		Where(sysuser.IDEQ(req.ID), sysuser.DeletedAtIsNil()).
		Only(ctx)
}

// GetSysUserByUserNameAll 用户名不可以重复 即使之前的数据被删除或者停用
func (r *sysUserRepo) GetSysUserByUserNameAll(ctx context.Context, req string) (*ent.SysUser, error) {
	return r.data.EntDb.SysUser.
		Query().
		Where(sysuser.UserNameEQ(req)).
		Only(ctx)
}

// GetSysUserByMobileAll 手机号码可以重复 如果之前的数据被删除或者停用
func (r *sysUserRepo) GetSysUserByMobileAll(ctx context.Context, req string) (*ent.SysUser, error) {
	return r.data.EntDb.SysUser.
		Query().
		Where(sysuser.MobileEQ(req), sysuser.StatusEQ(sysuser.Status0), sysuser.DeletedAtIsNil()).
		Only(ctx)
}

func (r *sysUserRepo) CreateSysUser(ctx context.Context, req model.CreateSysUserReq, username string) (*ent.SysUser, error) {
	hashPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	return r.data.EntDb.SysUser.
		Create().
		// 关联岗位信息
		AddUsersPostIDs(req.PostIds...).
		SetDeptID(req.DeptID).
		SetUserName(req.UserName).
		SetRealName(req.RealName).
		SetMobile(req.Mobile).
		SetPassword(hashPassword).
		SetSex(sysuser.Sex(req.Sex)).
		SetStatus(sysuser.Status0).
		SetRemark(req.Remark).
		SetCreatedAt(time.Now()).
		SetCreatedBy(username).
		Save(ctx)
}

func (r *sysUserRepo) UpdateSysUser(ctx context.Context, req model.UpdateSysUserReq, username string) (*ent.SysUser, error) {
	q := r.data.EntDb.SysUser.UpdateOneID(req.ID)
	// 先删除关联数据 再新增
	if len(req.PostIds) > 0 {
		q = q.ClearUsersPost().AddUsersPostIDs(req.PostIds...)
	}
	// 先清除历史关联
	return q.
		SetDeptID(req.DeptID).
		SetRealName(req.RealName).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetRemark(req.Remark).
		SetSex(sysuser.Sex(req.Sex)).
		SetStatus(sysuser.Status(req.Status)).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy(username).
		Save(ctx)
}

func (r *sysUserRepo) DeleteSysUser(ctx context.Context, req model.ByIdRequest, username string) (bool, error) {
	_, err := r.data.EntDb.SysUser.
		UpdateOneID(req.ID).
		SetDeletedAt(time.Now()).
		SetUpdatedBy(username).
		Save(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *sysUserRepo) PageSysUser(ctx context.Context, req model.PageSysUserReq) ([]*ent.SysUser, int, error) {
	q := r.data.EntDb.SysUser.Query()

	if len(req.Status) > 0 && (req.Status == "0" || req.Status == "1") {
		q = q.Where(sysuser.StatusEQ(sysuser.Status(req.Status)))
	}

	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// 排序方式
	var orderFn func(model.PageOrder) ent.OrderFunc
	orderFn = func(orderBy model.PageOrder) ent.OrderFunc {
		switch orderBy {
		case model.CreatedAtAsc:
			return ent.Asc(sysuser.FieldCreatedAt)
		case model.SysUserUsernameAsc:
			return ent.Asc(sysuser.FieldUserName)
		case model.SysUserUsernameDesc:
			return ent.Desc(sysuser.FieldUserName)
		default:
			return ent.Desc(sysuser.FieldCreatedAt)
		}
	}

	data, err := q.Order(orderFn(req.Order)).Offset(req.GetOffset()).Limit(req.GetLimit()).All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
