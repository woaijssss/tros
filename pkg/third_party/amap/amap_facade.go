package amap

import (
	"context"
	trlogger "github.com/woaijssss/tros/logx"
)

var Client = new(client)

type Location struct {
	LocationStr string  // 经纬度字符串，形如 113.538507,22.098750
	Longitude   float64 // 经度
	Latitude    float64 // 纬度
}

func (c *client) ParseLocation(ctx context.Context, name string) (*Location, error) {
	return parseLocation(ctx, name)
}

// RegeoBase 逆地理编码——基本信息（只需要经纬度时，推荐使用该接口）
func (c *client) RegeoBase(ctx context.Context, location *Location) (*regeoBaseResponse, error) {
	return c.regeoBase(ctx, location)
}

//// RegeoAll 逆地理编码——全部信息（适用于需要点位周边信息时调用）
//func (c *client) RegeoAll(ctx context.Context, location *Location) {
//	return c.regeoAll(ctx, location)
//}

type GetRoutePlanOption struct {
	Origin      string   // 出发点 经度在前，纬度在后，经度和纬度用","分割，经纬度小数点后不得超过6位。格式为x1,y1|x2,y2|x3,y3。   经纬度小数点不超过6位
	Destination string   // 目的地 经度在前，纬度在后，经度和纬度用","分割，经纬度小数点后不得超过6位。
	Strategy    string   // 驾车选择策略（暂不支持，todo 未来需要支持）
	Waypoints   []string // 途经点 经度和纬度用","分割，经度在前，纬度在后，小数点后不超过6位，坐标点之间用";"分隔；最大数目：16个坐标点。如果输入多个途径点，则按照用户输入的顺序进行路径规划
	// todo 后续增加车牌省份、车辆类型等参数，用于适配“北京”这种自驾特殊的城市
}

// GetRoutePlan 路径规划接口
func (c *client) GetRoutePlan(ctx context.Context, opt *GetRoutePlanOption) ([]string, error) {
	if opt == nil {
		return []string{}, nil
	}

	return c.getRoutePlan(ctx, opt)
}

type GetLiveWeatherOption struct {
	Province      string `json:"province"`      // 省份名
	City          string `json:"city"`          // 城市名
	AdCode        string `json:"adcode"`        // 区域编码
	Weather       string `json:"weather"`       // 天气现象（汉字描述）
	Temperature   string `json:"temperature"`   // 实时气温，单位：摄氏度
	WindDirection string `json:"winddirection"` // 风向描述
	WindPower     string `json:"windpower"`     // 风力级别，单位：级
	Humidity      string `json:"humidity"`      // 空气湿度
	ReportTime    string `json:"reporttime"`    // 数据发布的时间
}

// GetLiveWeather 实时天气
func (c *client) GetLiveWeather(ctx context.Context, adCode int64) (*GetLiveWeatherOption, error) {
	live, err := c.getLiveWeather(ctx, adCode)
	if err != nil {
		trlogger.Errorf(ctx, "GetLiveWeather err: [%+v]", err)
		return nil, err
	}

	return &GetLiveWeatherOption{
		Province:      live.Province,
		City:          live.City,
		AdCode:        live.AdCode,
		Weather:       live.Weather,
		Temperature:   live.Temperature,
		WindDirection: live.WindDirection,
		WindPower:     live.WindPower,
		Humidity:      live.Humidity,
		ReportTime:    live.ReportTime,
	}, nil
}

type SearchScenicByNameOption struct {
	Name string
}

type SearchScenicByNameResponse struct {
	Pois  []*Poi // 点位合集
	Total int32  // 总数
}

// SearchScenicByName 根据名字搜索景点
func (c *client) SearchScenicByName(ctx context.Context, opt *SearchScenicByNameOption) (*SearchScenicByNameResponse, error) {
	if opt == nil {
		return &SearchScenicByNameResponse{}, nil
	}

	return c.searchScenicByName(ctx, opt)
}
