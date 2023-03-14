package biz

import (
	"errors"

	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

var ProviderSetBiz = wire.NewSet(
	NewSysUserUseCase,
)

func structCopy(logger *zap.SugaredLogger, toVal, fromVal interface{}) error {
	err := copier.Copy(toVal, fromVal)
	if err != nil {
		logger.Errorf("结构体转换异常：%v", err)
		return errors.New("结构体信息异常")
	}
	return nil
}
