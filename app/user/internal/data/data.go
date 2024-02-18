package data

import (
	"realworld/app/user/internal/conf"
	rwredis "realworld/package/rwRedis"
	rwsql "realworld/package/rwSql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *gorm.DB
	redisCli *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (data *Data, function func(), err error) {
	user_db := NewDataDB(c)
	user_redis := NewRedisCli(c)

	function = func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	data = &Data{
		db:       user_db,
		redisCli: user_redis,
	}
	return
}

func NewDataDB(c *conf.Data) (db *gorm.DB) {
	db = rwsql.NewGormDb(&rwsql.GormConfig{
		Source: c.Database.Source,
	})
	return
}

func NewRedisCli(c *conf.Data) (redisCli *redis.Client) {
	redisCli = rwredis.NewRedis(&rwredis.RedisConfig{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		PoolSize:     int(c.Redis.PoolSize),
	})
	return
}
