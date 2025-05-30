package constants

import "time"

const (
	Profile    = "profile"
	TraceId    = "trace_id"
	UID        = "uid"
	OpId       = "op-id"
	RunAs      = "run-as"
	Roles      = "roles"
	BizTypes   = "biz-types"
	GroupId    = "group-id"
	Platform   = "platform"
	UserAgent  = "user_agent"
	Lang       = "lang"
	GoId       = "goid"
	PageNo     = "pageNo"
	PageSize   = "pageSize"
	Token      = "token"
	UserId     = "user_id"
	ShareToken = "s-token"
	RemoteIp   = "remote-ip"
	RequestUrl = "request-url"
	CompanyId  = "company-id"
	Product    = "product"
)

const GTimeout = time.Hour // 全局默认过期时间 1小时

type SystemEnv int32

const (
	SystemEnvUuid           = "0001" // 固定id，db索引
	Prod          SystemEnv = 1
	Test          SystemEnv = 2
)

const MaxApiExecuteTimeMs = 2000 // 接口请求默认最大时间
