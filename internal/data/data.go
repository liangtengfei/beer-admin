package data

import (
	"context"
	"database/sql"
	"time"

	ent2 "entgo.io/ent"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	"github.com/liangtengfei/beer-admin/internal/conf"
	"github.com/liangtengfei/beer-admin/internal/data/ent"
	"github.com/liangtengfei/beer-admin/internal/data/ent/migrate"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ProviderSetData = wire.NewSet(
	NewData,
	NewEntClient,
	NewRedisCmd,
	NewSysUserRepo,
)

type Data struct {
	EntDb *ent.Client
	Redis redis.Cmdable
}

func NewEntClient(conf *conf.Data, logger *zap.SugaredLogger) *ent.Client {
	log := logger.Named("module/data/ent")

	db, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv)).Debug()

	// 全局Hook
	client.Use(func(next ent2.Mutator) ent2.Mutator {
		return ent2.MutateFunc(func(ctx context.Context, m ent2.Mutation) (ent2.Value, error) {
			start := time.Now()
			defer func() {
				log.Infof("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()

			return next.Mutate(ctx, m)
		})
	})

	if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	return client
}

func NewRedisCmd(conf *conf.Data, logger *zap.SugaredLogger) redis.Cmdable {
	log := logger.Named("module/data/redis")
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  *conf.Redis.ReadTimeout,
		WriteTimeout: *conf.Redis.WriteTimeout,
		DialTimeout:  *conf.Redis.DialTimeout,
		PoolSize:     10,
		DB:           0,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis连接错误: %v", err)
	}

	return client
}

func NewData(entClient *ent.Client, redisCmd redis.Cmdable, logger *zap.SugaredLogger) (*Data, func(), error) {
	log := logger.Named("module/data")
	d := &Data{
		EntDb: entClient,
		Redis: redisCmd,
	}

	return d, func() {
		if err := d.EntDb.Close(); err != nil {
			log.Errorf("关闭数据库连接失败: %v", err)
		}
	}, nil
}
