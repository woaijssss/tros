package ip

import (
	"context"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	asnmap "github.com/projectdiscovery/asnmap/libs"
	trlogger "github.com/woaijssss/tros/logx"
	"net"
	"strconv"
	"strings"
)

// IPInfo 定义你需要的返回结构
type IPInfo struct {
	Country string `json:"country"` // 国家
	Region  string `json:"region"`  // 地区
	City    string `json:"city"`    // 城市
	Isp     string `json:"isp"`     // 运营商

	SubnetMask string `json:"subnet_mask"` // 子网掩码 (Windows/Linux通用)
	IPStart    string `json:"ip_start"`    // IP段起始
	IPEnd      string `json:"ip_end"`      // IP段结束

	// 位置信息
	Lon string `json:"lon"` // 经度
	Lat string `json:"lat"` // 纬度

	AsNumberInfo // AS号信息
}

type AsNumberInfo struct {
	Timestamp string
	Input     string
	ASN       string // 组织
	ASNOrg    string
	ASCountry string // 组织所在的国家
	ASRange   []string
}

func (c *client) GetIpRegion(ctx context.Context, ip string) (*IPInfo, error) {
	// 2. 测试查询
	info, err := c.queryIPInfo(ctx, ip)
	if err != nil {
		return nil, err
	}

	respAsNumberInfo, err := c.getAsNumber(ctx, ip)
	if err != nil {
		return nil, err
	}

	info.Timestamp = respAsNumberInfo.Timestamp
	info.Input = respAsNumberInfo.Input
	info.ASN = respAsNumberInfo.ASN
	info.ASNOrg = respAsNumberInfo.ASNOrg
	info.ASCountry = respAsNumberInfo.ASCountry
	info.ASRange = respAsNumberInfo.ASRange

	if c.geoIpDb != nil {
		cityGeoInfo, err := c.getIpLocation(ctx, ip)
		if err != nil {
			// pass
		} else {
			info.Lon = strconv.FormatFloat(cityGeoInfo.Location.Longitude, 'f', -1, 64)
			info.Lat = strconv.FormatFloat(cityGeoInfo.Location.Latitude, 'f', -1, 64)
		}
	}

	return info, nil
}

type getIpLocationResult struct {
}

func (c *client) getIpLocation(ctx context.Context, ipStr string) (*geoip2.City, error) {
	return c.geoIpDb.City(net.ParseIP(ipStr))
}

func (c *client) getAsNumber(ctx context.Context, ip string) (*AsNumberInfo, error) {
	cli, err := asnmap.NewClient()
	if err != nil {
		trlogger.Errorf(ctx, "getAsNumber asnmap.NewClient err: [%+v]", err)
		return nil, err
	}

	// 2. 构造查询请求（这里演示通过 IP 地址查询 ASN）
	// 你也可以换成 libs.ASNQuery("AS12345") 或 libs.OrgQuery("Google")
	responses, err := cli.GetData(ip)
	if err != nil {
		return nil, err
	}

	// 转换为易读的格式
	results, err := asnmap.MapToResults(responses)
	if err != nil {
		trlogger.Errorf(ctx, "getAsNumber asnmap.MapToResults err: [%+v]", err)
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("getAsNumber asnmap.MapToResults no results")
	}

	r := results[0]

	return &AsNumberInfo{
		Timestamp: r.Timestamp,
		Input:     r.Input,
		ASN:       r.ASN,
		ASNOrg:    r.ASN_org,
		ASCountry: r.AS_country,
		ASRange:   r.AS_range,
	}, nil
}

// 查询并解析 IP 信息
func (c *client) queryIPInfo(ctx context.Context, ipStr string) (*IPInfo, error) {
	// 调用底层库查询，返回格式如："中国|0|江苏省|南京市|电信"
	regionStr, err := c.searcher.SearchByStr(ipStr)
	if err != nil {
		trlogger.Errorf(ctx, "queryIPInfo c.searcher.SearchByStr err: [%+v]", err)
		return nil, err
	}

	trlogger.Infof(ctx, "queryIPInfo c.searcher.SearchByStr regionStr: [%s]", regionStr)
	// 解析归属地和运营商 (ip2region 格式：国家|区域|省份|城市|运营商)
	parts := strings.Split(regionStr, "|")
	country, region, city, isp := "", "", "", ""
	if len(parts) >= 5 {
		// 拼接国家、省、市作为归属地区
		//region = strings.Join([]string{parts[0], parts[2], parts[3]}, "")
		country = parts[0]
		region = parts[1]
		city = parts[2]
		isp = parts[3]
	}

	// 获取 IP 段起始、结束以及子网掩码
	ipStart, ipEnd, mask := getIPRangeAndMask(ipStr)

	return &IPInfo{
		Country:    country,
		Region:     region,
		City:       city,
		Isp:        isp,
		SubnetMask: mask,
		IPStart:    ipStart,
		IPEnd:      ipEnd,
	}, nil
}

// getIPRangeAndMask 模拟获取 IP 段及计算子网掩码
// 注：标准 ip2region 查询接口不直接返回该 IP 所属网段的起始/结束值。
// 在实际业务中，如果你需要绝对精准的数据库网段边界，通常需要修改底层库的查询方法返回对应的索引信息。
// 这里演示通过标准库计算常规子网掩码和网段。
func getIPRangeAndMask(ipStr string) (start, end, mask string) {
	ip := net.ParseIP(ipStr)
	fmt.Println("ip: ", ip)
	// 假设这是一个常见的 C 类子网 /24 (实际开发中可结合具体业务逻辑或自定义网段库)
	_, ipNet, _ := net.ParseCIDR(ipStr + "/24")

	// 计算子网掩码
	maskIP := net.IP(ipNet.Mask).String()

	// 计算网段起始 IP (网络地址)
	startIP := ipNet.IP.String()

	// 计算网段结束 IP (广播地址)
	endIP := make(net.IP, len(ipNet.IP))
	copy(endIP, ipNet.IP)
	for i := 0; i < len(ipNet.Mask); i++ {
		endIP[i] |= ^ipNet.Mask[i]
	}

	return startIP, endIP.String(), maskIP
}
