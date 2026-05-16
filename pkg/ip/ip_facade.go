package ip

import (
	"context"
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/oschwald/geoip2-golang"
	asnmap "github.com/projectdiscovery/asnmap/libs"
	trlogger "github.com/woaijssss/tros/logx"
	"sync"
)

type Client interface {
	Ip2Region(ctx context.Context, ipStr string) string
	Ip2RegionFullInfo(ctx context.Context, ipStr string) (*IPInfo, error)
}

type client struct {
	xdbPath    string
	searcher   *xdb.Searcher
	geoIpDb    *geoip2.Reader
	pdcpApiKey string
}

var (
	instance Client    // 单例实例
	once     sync.Once // 保证只执行一次的核心工具
	initErr  error     // 1. 声明一个全局的错误变量，用来记录初始化时的错误
)

// GetInstance 获取单例实例的全局方法（线程安全）- 仅ip信息
func GetInstance(ctx context.Context, xdbPath string, pdcpApiKey string) (Client, error) {
	once.Do(func() {
		// 这里的初始化逻辑，无论多少个协程同时调用 GetInstance，都只会执行一次
		cli, err := newClientWithSearcher(ctx, xdbPath, pdcpApiKey)
		if err != nil {
			initErr = err
			return
		}
		fmt.Println("GetInstance 正在初始化单例对象...")
		instance = cli
	})

	// 3. Do 执行完毕后，检查全局的 initErr 是否有值
	if initErr != nil {
		return nil, initErr
	}

	return instance, nil
}

// GetInstanceWithXdbAndGeoDb 获取单例实例的全局方法（线程安全）- 带经纬度数据库
func GetInstanceWithXdbAndGeoDb(ctx context.Context, xdbPath string, geoLite2MmdbPath string, pdcpApiKey string) (Client, error) {
	once.Do(func() {
		// 这里的初始化逻辑，无论多少个协程同时调用 GetInstance，都只会执行一次
		cli, err := newClientWithXdbAndGeoDb(ctx, xdbPath, geoLite2MmdbPath, pdcpApiKey)
		if err != nil {
			initErr = err
			return
		}
		fmt.Println("GetInstanceWithXdbAndGeoDb 正在初始化单例对象...")
		instance = cli
	})

	// 3. Do 执行完毕后，检查全局的 initErr 是否有值
	if initErr != nil {
		return nil, initErr
	}

	return instance, nil
}

func NewClient(xdbPath string) Client {
	return &client{
		xdbPath: xdbPath,
	}
}

func newClientWithSearcher(ctx context.Context, xdbPath string, pdcpApiKey string) (Client, error) {
	searcher, err := xdb.NewWithFileOnly(xdbPath)
	if err != nil {
		trlogger.Errorf(ctx, "NewClientWithSearcher init xdb err: [%+v]", err)
		return nil, err
	}

	asnmap.PDCPApiKey = pdcpApiKey // 设置pdcp key

	return &client{
		xdbPath:    xdbPath,
		pdcpApiKey: pdcpApiKey,
		searcher:   searcher,
	}, nil
}

func newClientWithXdbAndGeoDb(ctx context.Context, xdbPath string, geoLite2MmdbPath string, pdcpApiKey string) (Client, error) {
	searcher, err := xdb.NewWithFileOnly(xdbPath)
	if err != nil {
		trlogger.Errorf(ctx, "newClientWithXdbAndGeoDb init xdb err: [%+v]", err)
		return nil, err
	}

	geoIpDb, err := geoip2.Open(geoLite2MmdbPath)
	if err != nil {
		trlogger.Errorf(ctx, "newClientWithXdbAndGeoDb init geoIpDb err: [%+v]", err)
		return nil, err
	}

	asnmap.PDCPApiKey = pdcpApiKey // 设置pdcp key

	return &client{
		xdbPath:    xdbPath,
		pdcpApiKey: pdcpApiKey,
		searcher:   searcher,
		geoIpDb:    geoIpDb,
	}, nil
}

func (c *client) Ip2Region(ctx context.Context, ipStr string) string {
	region, err := c.searcher.SearchByStr(ipStr)
	if err != nil {
		trlogger.Errorf(ctx, "ip region search err: [%+v]", err)
		return ""
	}

	return region
}

func (c *client) Ip2RegionFullInfo(ctx context.Context, ipStr string) (*IPInfo, error) {
	return c.GetIpRegion(ctx, ipStr)
}
