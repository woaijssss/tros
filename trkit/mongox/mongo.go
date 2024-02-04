package mongox

import (
	"gitee.com/idigpower/tros/trerror"
)

var MongoxErrNoFound = trerror.DefaultTrError("MONGO NO FOUND")
var MongoxErrInsertError = trerror.DefaultTrError("MONGO INSERT ERROR")
var MongoxErrNoData = trerror.DefaultTrError("NO DATA")

type MongoConfig struct {
	MongoConfPassword string
	MongoConfUser     string
	MongoConfUrl      string
}

const MongoConfUserOnline = "root"
const MongoConfPasswordOnline = "123456"
const MongoConfUrlOnline = "127.0.0.1:28121"

const MongoConfUserDev = "root"
const MongoConfPasswordDev = "123456"
const MongoConfUrlDev = "127.0.0.1:28121"

func GetMongoConf(env string) (conf MongoConfig) {
	if env == "offline" {
		conf.MongoConfUser = MongoConfUserDev
		conf.MongoConfPassword = MongoConfPasswordDev
		conf.MongoConfUrl = MongoConfUrlDev

		return
	}

	conf.MongoConfUser = MongoConfUserOnline
	conf.MongoConfPassword = MongoConfPasswordOnline
	conf.MongoConfUrl = MongoConfUrlOnline
	return
}
