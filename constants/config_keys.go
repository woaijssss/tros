package constants

const (
	EnvConfigFile    = "config_file"
	GlobalConfigFile = "./conf/app.yaml"
)

const (
	AppId             = "app.app_id"
	AppName           = "app.app_name"
	AppVersion        = "app.app_version"
	GlobalEnv         = "app.env"
	GlobalHttpPort    = "app.httpPort"
	GlobalGrpcPort    = "app.grpcPort"
	GlobalMonitorPort = "app.monitorPort"
	LogLevel          = "app.log.level"
	LogPath           = "app.log.path"

	// MysqlUrl 数据库url
	MysqlUrl = "mysql.url"
	// MysqlPoolSize 最大连接数
	MysqlPoolSize = "mysql.poolSize"
	// MysqlMaxLife 连接的最大生命周期，单位是秒
	MysqlMaxLife = "mysql.maxLife"
	// MysqlMaxIdleCons 最大空闲连接数
	MysqlMaxIdleCons = "mysql.maxIdleCons"
	// MysqlMaxIdleTime 最大空闲时间，单位是秒
	MysqlMaxIdleTime = "mysql.maxIdleTime"
	// MysqlLog 该在该数据源上执行sql是是否需要把待执行的sql输出到日志
	MysqlLog = "mysql.log"
)
