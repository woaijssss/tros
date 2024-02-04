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

	MysqlUrl         = "mysql.url"
	MysqlPoolSize    = "mysql.poolSize"
	MysqlMaxLife     = "mysql.maxLife"
	MysqlMaxIdleCons = "mysql.maxIdleCons"
	MysqlMaxIdleTime = "mysql.maxIdleTime"
	MysqlLog         = "mysql.log"
)
