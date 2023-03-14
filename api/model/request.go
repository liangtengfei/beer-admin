package model

// PaginationRequest 分页通用请求参数
type PaginationRequest struct {
	PageNum   int    `json:"pageNum" form:"pageNum" binding:"required,min=1"`   // 页码
	PageSize  int    `json:"pageSize" form:"pageSize" binding:"required,min=1"` // 每页条数
	Status    string `json:"status" form:"status" binding:"omitempty,oneof=0 1"`
	SortField string `json:"sortField" form:"sortField"` // 排序字段
	SortOrder string `json:"sortOrder" form:"sortOrder"` // 排序顺序
}

type ByIdRequest struct {
	ID uint64 `json:"id" form:"id" uri:"id" binding:"number,min=0"`
}

type ByIdStrRequest struct {
	ID string `json:"id" form:"id" uri:"id" binding:"required,uuid"`
}

type ByIdsRequest struct {
	ID []uint64 `json:"id" form:"id" uri:"id" binding:"required,gt=0"`
}

type StatusUpdateRequest struct {
	ID     uint64 `json:"id" binding:"required"`               // 部门id
	Status string `json:"status" binding:"required,oneof=0 1"` // 状态（0正常 1停用）
}

func (req *PaginationRequest) GetOffset() int {
	if req.PageNum > 10000 || req.PageNum <= 0 {
		req.PageSize = 1
	}
	var offset int
	if req.PageSize <= 0 || req.PageSize > 10000 {
		offset = 10
	} else {
		offset = (req.PageNum - 1) * req.PageSize
	}
	return offset
}

func (req *PaginationRequest) GetLimit() int {
	var limit int
	if req.PageSize <= 0 || req.PageSize > 10000 {
		limit = 10
	} else {
		limit = req.PageSize
	}
	return limit
}

type PageOrder int32

const (
	CreatedAtDesc PageOrder = 0
	CreatedAtAsc  PageOrder = 1
)
