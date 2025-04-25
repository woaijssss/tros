package redisx

import (
	"context"
	trlogger "github.com/woaijssss/tros/logx"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	RedisClient   *redis.Pool
	RedisHost     string
	RedisPassword string
	Env           string
	RedisDB       int
)

func Setup(ctx context.Context) {
	initRedis(ctx)
}

func RedisSetup(ctx context.Context) {
	initRedis(ctx)
}

func SetupByHost(host, password string) {
	initRedisByHost(host, password)
}

func InitRedisByConf(conf *RedisConfig, env string) {
	if conf.RedisHost == "" {
		SetDefaultRedisConf(conf, env)
	}

	RedisHost = conf.RedisHost
	RedisDB = RedisDB
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     RedisPoolMaxIdle,
		MaxActive:   RedisPoolMaxActive,
		IdleTimeout: RedisPoolIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conf.RedisPassword); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

func initRedis(ctx context.Context) {
	//Env = env
	//conf := GetRedisConf(env)
	conf := GetRedisConfV2()
	// 从配置文件获取redis的ip以及db
	RedisHost = conf.RedisHost
	RedisDB = RedisDB
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     RedisPoolMaxIdle,
		MaxActive:   RedisPoolMaxActive,
		IdleTimeout: RedisPoolIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conf.RedisPassword); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

func initRedisByHost(host, password string) {
	// 从配置文件获取redis的ip以及db
	RedisHost = host
	RedisPassword = password
	RedisDB = RedisDB
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     RedisPoolMaxIdle,
		MaxActive:   RedisPoolMaxActive,
		IdleTimeout: RedisPoolIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", RedisPassword); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

func getRedisConn(ctx context.Context) redis.Conn {
	client := RedisClient.Get()
	if client.Err() != nil {
		trlogger.Fatalf(ctx, "Get redis string err %+v", client.Err())
		//initRedis(Env)
		initRedisByHost(RedisHost, RedisPassword)
		return RedisClient.Get()
	}
	return client
}

func ListAll(context context.Context, key string) ([]interface{}, error) {
	client := getRedisConn(context)

	defer client.Close()
	if client.Err() != nil {
		trlogger.Fatalf(context, "Get redis string err %+v", client.Err())
		return nil, client.Err()
	}

	value, err := redis.Values(client.Do("lrange", key, 0, -1))
	if err == redis.ErrNil {
		trlogger.Debugf(context, "Get redis string empty [%+v]", err)
		return nil, nil
	}

	if err != nil {
		trlogger.Fatalf(context, "Get redis string err %+v", err)
		return nil, err
	}

	return value, nil
}

func Get(context context.Context, key string) (string, error) {
	client := getRedisConn(context)

	defer client.Close()
	if client.Err() != nil {
		trlogger.Fatalf(context, "Get redis string err %+v", client.Err())
		return "", client.Err()
	}

	value, err := redis.String(client.Do("get", key))
	if err == redis.ErrNil {
		trlogger.Debugf(context, "Get redis string empty [%+v]", err)
		return "", nil
	}

	if err != nil {
		trlogger.Fatalf(context, "Get redis string err %+v", err)
		return "", err
	}

	return value, nil
}

func Set(ctx context.Context, key, val string, expire int64) error {
	client := getRedisConn(ctx)
	defer client.Close()

	if client.Err() != nil {
		trlogger.Fatalf(ctx, "Get redis string err %+v", client.Err())
		return client.Err()
	}

	_, err := client.Do("set", key, val, "ex", expire)
	if err != nil {
		trlogger.Fatalf(ctx, "Set redis string err %+v", err)
		return err
	}
	return err
}

func KeysAll(ctx context.Context, keyPrefix string) ([]string, error) {
	keyPrefix += "*" // 默认补充一个*

	client := getRedisConn(ctx)
	defer client.Close()

	var resultKeyList []string
	// 执行 KEYS 命令，获取匹配模式的键
	keys, err := redis.Strings(client.Do("KEYS", keyPrefix))
	if err != nil {
		trlogger.Fatalf(ctx, "keys all redis string err %+v", err)
		return []string{}, err
	}

	// 打印获取到的键
	for _, key := range keys {
		resultKeyList = append(resultKeyList, key)
	}
	return resultKeyList, nil
}

func Incr(context context.Context, key string) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		trlogger.Fatalf(context, "incr redis string err %+v", client.Err())
		return client.Err()
	}
	_, err := client.Do("incr", key)
	if err != nil {
		trlogger.Fatalf(context, "incr redis string err %+v", err)
		return err
	}
	return err
}

func SetV2(context context.Context, key, val string) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		trlogger.Fatalf(context, "Get redis string err %+v", client.Err())
		return client.Err()
	}

	_, err := client.Do("set", key, val)
	if err != nil {
		trlogger.Fatalf(context, "Set redis string err %+v", err)
		return err
	}
	return err
}

func Expire(context context.Context, key string, expire int64) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		trlogger.Fatalf(context, "expire redis string err %+v", client.Err())
		return client.Err()
	}
	_, err := client.Do("expire", key, expire)
	if err != nil {
		trlogger.Fatalf(context, "expire redis string err %+v", err)
		return err
	}
	return err
}

func Delete(context context.Context, key string) error {
	client := RedisClient.Get()
	defer client.Close()

	if client.Err() != nil {
		trlogger.Fatalf(context, "Get redis string err %+v", client.Err())
		return client.Err()
	}

	_, err := client.Do("DEL", key)
	if err != nil {
		trlogger.Fatalf(context, "redis delelte failed: %+v", err)
	}
	return err
}
func GetBit(context context.Context, key string, offset int64) (int64, error) {
	client := getRedisConn(context)

	defer client.Close()
	if client.Err() != nil {
		trlogger.Fatalf(context, "GetBit redis Err[%v]", client.Err())
		return 0, client.Err()
	}

	value, err := redis.Int64(client.Do("getbit", key, offset))
	if err == redis.ErrNil {
		trlogger.Debugf(context, "GetBit redis Err[%v]", err)
		return 0, nil
	}

	if err != nil {
		trlogger.Fatalf(context, "GetBit redis Err[%v]", err)
		return 0, err
	}

	return value, nil
}

func SetBit(context context.Context, key string, val, offset int64) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		trlogger.Fatalf(context, "SetBit Redis Err[%v]", client.Err())
		return client.Err()
	}

	_, err := client.Do("setbit", key, offset, val)
	if err != nil {
		trlogger.Fatalf(context, "SetBit Redis Err[%v]", err)
		return err
	}
	return err
}
