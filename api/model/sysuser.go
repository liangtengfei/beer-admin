package model

import (
	"time"

	"github.com/liangtengfei/beer-admin/pkg/datetime"
)

type SysUserInfo struct {
	// 系统用户
	ID uint64 `json:"userID"`
	// 创建时间
	CreatedAtStr string `json:"createdAt"`
	// 更新时间
	UpdatedAtStr string `json:"updatedAt"`
	// 创建人员
	CreatedBy string `json:"createdBy"`
	// 更新人员
	UpdatedBy string `json:"updatedBy"`
	// 部门编号
	DeptID uint64 `json:"deptID"`
	// 登陆用户名称
	UserName string `json:"userName"`
	// 真实姓名
	RealName string `json:"realName"`
	// 手机号码
	Mobile string `json:"mobile"`
	// 邮箱地址
	Email string `json:"email"`
	// 头像地址
	Avatar string `json:"avatar"`
	// 性别
	Sex string `json:"sex"`
	// 状态: 0=ON 1=OFF
	Status string `json:"status"`
	// 备注
	Remark string `json:"remark"`

	RoleIds []uint64 `json:"roleIds"`
	PostIds []uint64 `json:"postIds"`
}

func (res *SysUserInfo) CreatedAt(t time.Time) {
	res.CreatedAtStr = datetime.Full(t)
}

func (res *SysUserInfo) UpdatedAt(t time.Time) {
	res.UpdatedAtStr = datetime.Full(t)
}

func (res *SysUserInfo) IsAdmin() bool {
	return res.ID == uint64(1)
}

type RestSysUserPasswordReq struct {
	ID             uint64 `json:"id" binding:"required"`
	Password       string `json:"password" binding:"required,min=6,max=16,eqfield=RepeatPassword" minLength:"6" maxLength:"16"`
	RepeatPassword string `json:"repeatPassword" binding:"required,min=6,max=16,eqfield=Password" minLength:"6" maxLength:"16"`
}

type UpdateSysUserPasswordReq struct {
	ID             uint64 `json:"id" binding:"required"`
	Password       string `json:"password" binding:"required,min=6,max=16" minLength:"6" maxLength:"16"`
	NewPassword    string `json:"newPassword" binding:"required,min=6,max=16,eqfield=RepeatPassword" minLength:"6" maxLength:"16"`
	RepeatPassword string `json:"repeatPassword" binding:"required,min=6,max=16,eqfield=NewPassword" minLength:"6" maxLength:"16"`
}

type GetSysUserByUserNameReq struct {
	UserName string `json:"userName" binding:"required,alphanum"` // 用户账号
}

type CreateSysUserReq struct {
	UserName string   `json:"userName" binding:"required,alphanum"`                                  // 用户账号
	RealName string   `json:"realName" binding:"required"`                                           // 用户昵称
	Email    string   `json:"email" binding:"omitempty,email"`                                       // 用户邮箱
	DeptID   uint64   `json:"deptID" binding:"required"`                                             // 部门ID
	Remark   string   `json:"remark" binding:"-"`                                                    // 备注
	Password string   `json:"password" binding:"required,min=6,max=16" minLength:"6" maxLength:"16"` // 密码
	Avatar   string   `json:"avatar" binding:"-"`                                                    // 头像地址
	Sex      string   `json:"sex" binding:"omitempty,oneof=male female" enums:"0, 1"`                // 用户性别（0男 1女 2未知）
	Mobile   string   `json:"mobile" binding:"required"`                                             // 手机号码
	RoleIds  []uint64 `json:"roleIds" binding:"omitempty,dive,number" example:"1,2,3"`               // 角色编号数组
	PostIds  []uint64 `json:"postIds" binding:"-" example:"1,2,3"`                                   // 岗位编号数组
}

type UpdateSysUserReq struct {
	ID       uint64   `json:"id" binding:"required"`                                   // 用户ID
	DeptID   uint64   `json:"deptID" binding:"required"`                               // 部门ID
	RealName string   `json:"realName" binding:"required"`                             // 用户昵称
	Mobile   string   `json:"mobile" binding:"required"`                               // 手机号码
	Email    string   `json:"email" binding:"omitempty,email"`                         // 用户邮箱
	Sex      string   `json:"sex" binding:"omitempty,oneof=male female" enums:"0, 1"`  // 用户性别（0男 1女 2未知）
	Status   string   `json:"status" binding:"oneof=0 1"`                              // 状态（0正常 1停用）
	Remark   string   `json:"remark" binding:"-"`                                      // 备注
	RoleIds  []uint64 `json:"roleIds" binding:"omitempty,dive,number" example:"1,2,3"` // 角色编号数组
	PostIds  []uint64 `json:"postIds" binding:"-" example:"1,2,3"`                     // 岗位编号数组
}

type ListSysUserReq struct {
	DeptIDs []uint64 `json:"deptIds[]" form:"deptIds[]"` // 部门编号
}

const (
	SysUserUsernameAsc  PageOrder = 2
	SysUserUsernameDesc PageOrder = 3
)

type PageSysUserReq struct {
	PaginationRequest
	Order PageOrder `json:"order" form:"order" binding:"oneof=0 1 2 3"`
}
