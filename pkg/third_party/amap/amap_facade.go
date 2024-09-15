package amap

import (
	"context"
	trlogger "gitee.com/idigpower/tros/logx"
	"gitee.com/idigpower/tros/pkg/utils"
)

var Client = new(client)

type PlaceInfoV5 struct {
	PoiId    string // poi id
	Name     string // poi名称
	Location *Location
	Type     string // poi所属类型
	PName    string // poi所属省份
	CityName string // poi所属城市
	AdName   string // poi所属区县
	Address  string // poi详细地址

	Cost     string // 人均消费
	OpenTime string // 营业时间描述
	Rating   string // 评分

	Images []string // 地点相关图片
}
type Location struct {
	LocationStr string  // 经纬度字符串，形如 113.538507,22.098750
	Longitude   float64 // 经度
	Latitude    float64 // 纬度
}

func (c *client) GetPlaceDetail(ctx context.Context, name string) (string, *PlaceInfoV5, error) {
	trlogger.Infof(ctx, "GetPlaceDetail place name: [%s]", name)
	// poi查询
	response, err := c.getPlaceByName(ctx, name)
	if err != nil {
		trlogger.Errorf(ctx, "getPlaceByName err: [%+v]", err)
		return "", nil, err
	}

	// 解析poi结果
	rawData, placeInfoV5, err := c.parsePoiResult(ctx, response)
	if err != nil {
		return "", nil, err
	}

	// 格式化原始数据
	rawDataStr, err := utils.ToJsonString(rawData)
	if err != nil {
		trlogger.Errorf(ctx, "utils.ToJsonString PlaceInfoV5 raw data err: [%+v]", err)
		return "", placeInfoV5, err
	}
	trlogger.Infof(ctx, "GetPlaceDetail place by name success")
	return rawDataStr, placeInfoV5, err
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
