package ip

import (
	"context"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	trlogger "github.com/woaijssss/tros/logx"
)

type Client interface {
	Ip2Region(ctx context.Context, ipStr string) string
}

func NewClient(xdbPath string) Client {
	return &client{
		xdbPath: xdbPath,
	}
}

type client struct {
	xdbPath string
}

func (c *client) Ip2Region(ctx context.Context, ipStr string) string {
	// 初始化查询对象
	searcher, err := xdb.NewWithFileOnly(c.xdbPath)
	if err != nil {
		trlogger.Errorf(ctx, "ip region search init err: [%+v]", err)
		return ""
	}
	defer searcher.Close()

	region, err := searcher.SearchByStr(ipStr)
	if err != nil {
		trlogger.Errorf(ctx, "ip region search err: [%+v]", err)
		return ""
	}

	return region
}
