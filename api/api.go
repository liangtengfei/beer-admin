package api

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/liangtengfei/beer-admin/pkg/valid"
)

var NewSetAPI = wire.NewSet(
	SysUserSet,
)

type RestRes struct {
	Code   int         `json:"code" example:"200"`
	Msg    string      `json:"msg" example:"status ok"`
	Data   interface{} `json:"data,omitempty"`
	Status string      `json:"type" example:"success"`
}

type PaginationRes struct {
	Total int64 `json:"total"`
	RestRes
}

func response(code int, msg string, data interface{}) RestRes {
	res := RestRes{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	res.Status = _GetType(res.Code)
	return res
}

func Success(ctx *gin.Context) {
	res := response(http.StatusOK, "操作成功", gin.H{})
	ctx.JSON(http.StatusOK, res)
}

func SuccessData(ctx *gin.Context, data interface{}) {
	res := response(http.StatusOK, "操作成功", data)
	ctx.JSON(http.StatusOK, res)
}

func SuccessPage(ctx *gin.Context, total int, data interface{}) {
	res := PaginationRes{
		Total: int64(total),
		RestRes: RestRes{
			Code: http.StatusOK,
			Msg:  "操作成功",
			Data: gin.H{
				"list":  data,
				"total": total,
			},
			Status: "success",
		},
	}

	ctx.JSON(http.StatusOK, res)
}

// ErrorValid 请求参数验证错误时
func ErrorValid(ctx *gin.Context, err error) {
	var buff bytes.Buffer

	//避免格式错误
	switch err.(type) {
	case validator.ValidationErrors:
		m := valid.Translate(err)
		msgs := make([]string, 0)
		for _, v := range m {
			msgs = append(msgs, v...)
		}
		buff.WriteString(strings.Join(msgs, ","))

		ctx.AbortWithStatusJSON(http.StatusOK, response(http.StatusBadRequest, buff.String(), nil))
	default:
		ctx.AbortWithStatusJSON(http.StatusOK, response(http.StatusBadRequest, "参数格式/类型不正确", nil))
	}
}

func Error(ctx *gin.Context) {
	res := response(http.StatusInternalServerError, "操作错误", gin.H{})
	ctx.AbortWithStatusJSON(http.StatusOK, res)
}

func ErrorErr(ctx *gin.Context, err error) {
	res := response(http.StatusInternalServerError, err.Error(), gin.H{})
	ctx.AbortWithStatusJSON(http.StatusOK, res)
}

func ErrorMsg(ctx *gin.Context, msg string) {
	res := response(http.StatusInternalServerError, msg, gin.H{})
	ctx.AbortWithStatusJSON(http.StatusOK, res)
}

func _GetType(code int) string {
	if code == 200 {
		return "success"
	} else if code == -1 {
		return "warning"
	} else {
		return "error"
	}
}
