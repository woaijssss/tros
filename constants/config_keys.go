package constants

const (
	EnvConfigFile    = "config_file"
	GlobalConfigFile = "./conf/app.yaml"
)

const (
	AppId             = "app.app_id"      // 应用id
	AppName           = "app.app_name"    // 应用名字
	AppVersion        = "app.app_version" // 应用版本
	GlobalEnv         = "app.env"         // 应用运行环境
	GlobalHttpPort    = "app.httpPort"    // 监听http端口
	GlobalGrpcPort    = "app.grpcPort"    // 监听grpc端口
	GlobalMonitorPort = "app.monitorPort" // 监控Prometheus端口
	LogLevel          = "app.log.level"   // 日志级别
	LogPath           = "app.log.path"    // 日志路径

	MysqlUrl         = "mysql.url"         // 数据库url
	MysqlPoolSize    = "mysql.poolSize"    // 最大连接数
	MysqlMaxLife     = "mysql.maxLife"     // 连接的最大生命周期，单位是秒
	MysqlMaxIdleCons = "mysql.maxIdleCons" // 最大空闲连接数
	MysqlMaxIdleTime = "mysql.maxIdleTime" // 最大空闲时间，单位是秒
	MysqlLog         = "mysql.log"         // 该在该数据源上执行sql是是否需要把待执行的sql输出到日志

	IFlyTekId        = "iflytek.app_id"     // 科大讯飞api的appid
	IFlyTekSecretKey = "iflytek.secret_key" // 科大讯飞api的secret

	WechatAppid     = "wechat.app_id"     // 微信小程序的appid
	WechatAppSecret = "wechat.app_secret" // 微信小程序的app_secret
	WechatMchId     = "wechat.mch_id"     // 微信商户id
	WechatApiV2Key  = "wechat.apiv2_key"  // 微信支付的ApiV2密钥
	WechatApiV3Key  = "wechat.apiv3_key"  // 微信支付的ApiV3密钥

	AliOssBucket          = "oss.bucket"          // 阿里oss存储桶名字
	AliOssUrl             = "oss.url"             // 阿里oss存储桶访问主地址
	AliOssAccessKeyId     = "oss.accessKeyId"     // 阿里oss存储桶access_key
	AliOssAccessKeySecret = "oss.accessKeySecret" // 阿里oss存储桶access_secret
	AliOssBucketUrlPrefix = "oss.bucketUrlPrefix" // 阿里oss存储桶内容的访问前缀

	AMapAppKey = "amap.appKey" // 高德地图的appkey

	FeiShuWebHookUrlKey = "feishu.web_hook_url" // 飞书机器人webhook地址
	FeiShuSignKey       = "feishu.sign_key"     // 飞书通知签名key
)
