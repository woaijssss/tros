package mysqlx

import (
	"context"
	"gitee.com/idigpower/tros/conf"
	trlogger "gitee.com/idigpower/tros/logx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rolandhe/daog"
)

var globalDatasource daog.Datasource

var dbConf daog.DbConf

var valid string = "invalid"

func initResources(ctx context.Context) {
	if dbUrl := conf.GetMysqlUrl(); dbUrl != "" {
		var err error
		dbConf = daog.DbConf{
			DbUrl:    dbUrl,
			Size:     conf.GetMysqlPoolSize(),
			Life:     conf.GetMysqlMaxLife(),
			IdleCons: conf.GetMysqlMaxIdleCons(),
			IdleTime: conf.GetMysqlMaxIdleTime(),
			LogSQL:   true,
		}
		globalDatasource, err = daog.NewDatasource(&dbConf)
		if err != nil {
			panic(err)
		}
		valid = "valid"
	}
	trlogger.Infof(ctx, "mysql db is %s, conf is: [%+v]", valid, dbConf)
}
