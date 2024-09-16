package mysqlx

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rolandhe/daog"
	"github.com/woaijssss/tros/conf"
	trlogger "github.com/woaijssss/tros/logx"
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
