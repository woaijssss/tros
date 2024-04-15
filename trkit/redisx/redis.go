package redisx

import "gitee.com/idigpower/tros/conf"

const (
	HostDev     = "127.0.0.1:6379"
	PasswordDev = "123456"

	HostProd     = "127.0.0.1:6379"
	PasswordProd = "123456"
)

const (
	RedisPoolMaxIdle     = 100
	RedisPoolMaxActive   = 10000
	RedisPoolIdleTimeout = 1
)

const ExpireTime = 300

const (
	EnvOffline = "offline"
	EnvOnline  = "online"
)

type RedisConfig struct {
	RedisHost            string
	RedisPassword        string
	RedisPoolMaxIdle     int
	RedisPoolMaxActive   int
	RedisPoolIdleTimeout int
}

func GetRedisConf(env string) (conf RedisConfig) {
	if env == EnvOffline {
		conf.RedisHost = HostDev
		conf.RedisPassword = PasswordDev
		return
	}
	conf.RedisHost = HostProd
	conf.RedisPassword = PasswordProd
	return
}

func GetRedisConfV2() (redisConf RedisConfig) {
	//if env == EnvOffline {
	//	conf.RedisHost = HostDev
	//	conf.RedisPassword = PasswordDev
	//	return
	//}
	redisConf.RedisHost = conf.GetString("redis.host")
	redisConf.RedisPassword = conf.GetString("redis.password")
	return
}

func SetDefaultRedisConf(config *RedisConfig, env string) {
	if env == EnvOffline {
		config.RedisHost = HostDev
		config.RedisPassword = PasswordDev
	} else {
		config.RedisHost = HostProd
		config.RedisPassword = PasswordProd
	}

	if config.RedisHost == "" {
		if env == EnvOffline {
			config.RedisHost = HostDev
		} else {
			config.RedisHost = HostProd
		}
	}

	if config.RedisPassword == "" {
		if env == EnvOffline {
			config.RedisPassword = PasswordDev
		} else {
			config.RedisPassword = PasswordProd
		}
	}

	if config.RedisPoolIdleTimeout == 0 {
		config.RedisPoolIdleTimeout = RedisPoolIdleTimeout
	}

	if config.RedisPoolMaxActive == 0 {
		config.RedisPoolMaxActive = RedisPoolMaxActive
	}

	if config.RedisPoolMaxIdle == 0 {
		config.RedisPoolMaxIdle = RedisPoolMaxIdle
	}
}
