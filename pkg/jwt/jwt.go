package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/liangtengfei/beer-admin/api/model"
	"github.com/redis/go-redis/v9"
)

const (
	payloadKey = "auth_payload"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expire_at"`
}

func NewPayload(username string, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return jwtv4.ErrTokenExpired
	}
	return nil
}

// FromContext extract auth info from context
func FromContext(ctx *gin.Context) (token *Payload, ok bool) {
	token, ok = ctx.Value(payloadKey).(*Payload)
	return
}

func LoginUser(ctx *gin.Context, cache redis.Cmdable) (model.SysUserInfo, error) {
	var result model.SysUserInfo

	username, err := LoginUserName(ctx)
	if err != nil {
		return model.SysUserInfo{}, err
	}

	data, err := cache.Get(ctx, "sys:users:"+username).Result()
	if err != nil {
		return model.SysUserInfo{}, errors.New("未查询到登录信息")
	}

	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, fmt.Errorf("解析用户缓存失败：%s", err.Error())
	}

	return result, nil
}

func LoginUserName(ctx *gin.Context) (string, error) {
	token, ok := FromContext(ctx)
	if !ok {
		return "", errors.New("未登录")
	}
	return token.Username, nil
}
