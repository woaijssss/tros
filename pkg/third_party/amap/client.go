package amap

import (
	"context"
	"errors"
	"fmt"
	http2 "gitee.com/idigpower/tros/client/http"
	"gitee.com/idigpower/tros/conf"
	"gitee.com/idigpower/tros/constants"
	trlogger "gitee.com/idigpower/tros/logx"
	"strings"
)

type client struct {
}

const (
	placeDetailUrlTemplateV3 = "https://restapi.amap.com/v3/place/text?keywords=%s&city=beijing&offset=1&page=1&key=%s&extensions=all"
	placeDetailUrlTemplateV5 = "https://restapi.amap.com/v5/place/text?keywords=%s&offset=1&page_size=1&page_num=1&key=%s&extensions=all&show_fields=business,photos"

	// 逆地理编码接口地址
	/// 返回基本地址信息
	regeoExtensionsBaseUrl = "https://restapi.amap.com/v3/geocode/regeo?key=%s&location=%s&extensions=base"
	/// 返回全部信息
	regeoExtensionsAllUrl = "https://restapi.amap.com/v3/geocode/regeo?key=%s&location=%s&extensions=base"

	// 路径规划接口地址
	routePlanV1 = "https://restapi.amap.com/v3/direction/driving?key=%s&origin=%s&destination=%s&waypoints=%s"

	// 实时天气接口
	weatherLive = "https://restapi.amap.com/v3/weather/weatherInfo?extensions=base&output=JSON&key=%s&city=%d"
)

func (c *client) getPlaceByName(ctx context.Context, name string) (*http2.Response, error) {
	url := fmt.Sprintf(placeDetailUrlTemplateV5, name, conf.Get(constants.AMapAppKey))
	httpClient := http2.NewHttpClient()
	httpClient.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	return httpClient.Get(ctx, url)
}

// commonResponse 高德返回的基本结构（所有接口都需要此参数）
type commonResponse struct {
	Status string `json:"status"` // 本次API访问状态，如果成功返回1，如果失败返回0。
	Info   string `json:"info"`
}

// poi 2.0定义: https://lbs.amap.com/api/webservice/guide/api-advanced/newpoisearch
type response struct {
	Status string `json:"status"` // 本次API访问状态，如果成功返回1，如果失败返回0。
	Info   string `json:"info"`
	Pois   []*poi `json:"pois"`
}

type pois struct { // poi完整集合
	Poi      *poi      `json:"poi"`
	Children *children `json:"children"`
	Business *business `json:"business"`
	Indoor   *indoor   `json:"indoor"`
	Navi     *navi     `json:"navi"`
	Photos   []*photos `json:"photos"`
}

type poi struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Location string `json:"location"`
	Type     string `json:"type"`
	TypeCode string `json:"typecode"`
	PName    string `json:"pname"`
	CityName string `json:"cityname"`
	AdName   string `json:"adname"`
	Address  string `json:"address"`
	PCode    string `json:"pcode"`
	AdCode   string `json:"adcode"`
	CityCode string `json:"citycode"`

	Children *children `json:"children"`
	Business *business `json:"business"`
	Indoor   *indoor   `json:"indoor"`
	Navi     *navi     `json:"navi"`
	Photos   []*photos `json:"photos"`
}
type children struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Address  string `json:"address"`
	SubType  string `json:"subtype"`
	TypeCode string `json:"typecode"`
}
type business struct {
	BusinessArea  string `json:"business_area"`
	OpenTimeToday string `json:"opentime_today"`
	OpenTimeWeek  string `json:"opentime_week"`
	Tel           string `json:"tel"`
	Tag           string `json:"tag"`
	Rating        string `json:"rating"`
	Cost          string `json:"cost"`
	ParkingType   string `json:"parking_type"`
	Alias         string `json:"alias"`
}
type indoor struct {
	IndoorMap string `json:"indoor_map"`
	CpId      string `json:"cpid"`
	Floor     string `json:"floor"`
	TrueFloor string `json:"truefloor"`
}
type navi struct {
	NaviPoiId    string `json:"navi_poiid"`
	EntrLocation string `json:"entr_location"`
	ExitLocation string `json:"exit_location"`
	GridCode     string `json:"gridcode"`
}
type photos struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func (c *client) parsePoiResult(ctx context.Context, resp *http2.Response) (*response, *PlaceInfoV5, error) {
	var err error
	//// 从响应体中读取数据
	//respBody, _ := ioutil.ReadAll(resp.Body)
	//// 将JSON格式的响应转换为map类型
	//var result map[string]interface{}
	//result, err := utils.ByteToJson[map[string]interface{}](respBody)
	//if err != nil {
	//	trlogger.Errorf(ctx, "http response ByteToJson err: [%+v]", err)
	//	return nil, err
	//}
	//fmt.Println(result)

	var data response
	err = http2.ResToObj(resp, &data)
	if err != nil {
		trlogger.Errorf(ctx, "parsePoiResult http2.ResToObj err: [%+v]", err)
		return nil, nil, err
	}
	trlogger.Errorf(ctx, "parsePoiResult response: [%+v]", data)
	if len(data.Pois) != 1 {
		return &data, nil, fmt.Errorf("parsePoiResult poi not 1, [%d]", len(data.Pois))
	}
	// 输出查询到的详情信息
	//pois = data.Pois
	//poiObjs := pois[0].(map[string]interface{})
	//business := pois[0].(map[string]interface{})["business"].(map[string]interface{})
	//photos := pois[0].(map[string]interface{})["photos"].([]interface{})
	r := &PlaceInfoV5{
		PoiId:    data.Pois[0].Id,
		Name:     data.Pois[0].Name,
		Type:     data.Pois[0].Type,
		PName:    data.Pois[0].PName,
		CityName: data.Pois[0].CityName,
		AdName:   data.Pois[0].AdName,
		Address:  data.Pois[0].Address,

		Cost:     data.Pois[0].Business.Cost,
		OpenTime: data.Pois[0].Business.OpenTimeWeek,
		Rating:   data.Pois[0].Business.Rating,
	}

	l, err := parseLocation(ctx, data.Pois[0].Location)
	if err != nil {
		trlogger.Errorf(ctx, "parseLocation err: [%+v]", err)
		return &data, nil, err
	}
	r.Location = l

	for _, v := range data.Pois[0].Photos {
		r.Images = append(r.Images, v.Url)
	}

	return &data, r, nil
}

// ///////////////////////////////////////////////////// 路径规划结构
type getRoutePlanResponse struct {
	Status string `json:"status"`
	Info   string `json:"info"`
	Route  *route `json:"route"`
}

type route struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	TaxiCost    string  `json:"taxi_cost"` // 打车费用	单位：元，注意：extensions=all时才会返回
	Paths       []*path `json:"paths"`
}

// 驾车换乘方案
type path struct {
	Steps []*step `json:"steps"`
}

// 导航路段
type step struct {
	Polyline string `json:"polyline"` // 此路段坐标点串 格式为坐标串，如：116.481247,39.990704;116.481270,39.990726
}

/////////////////////////////////////////////////////// 路径规划结构end

// ///////////////////////////////////////////////////// 实时天气结构
type weatherLiveResponse struct {
	Status string             `json:"status"`
	Info   string             `json:"info"`
	Lives  []*weatherLiveBase `json:"lives"`
}

type weatherLiveBase struct {
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

// ///////////////////////////////////////////////////// 实时天气结构end

// ///////////////////////////////////////////////////// 逆地理编码结构
type regeoBaseResponse struct {
	Status    string     `json:"status"` // 本次API访问状态，如果成功返回1，如果失败返回0。
	Info      string     `json:"info"`
	RegeoCode *regeoCode `json:"regeocode"` // 逆地理编码列表
}

type regeoCode struct {
	AddressComponent *addressComponent `json:"addressComponent"` // 地址元素列表
}

type addressComponent struct {
	Country  string `json:"country"`  // 国家名称
	Province string `json:"province"` // 省份名称
	//City             string `json:"city"`              // 城市名称
	Citycode         string `json:"citycode"`          // 城市编码
	District         string `json:"district"`          // 所在地区名称
	Adcode           string `json:"adcode"`            // 行政区编码
	Township         string `json:"township"`          // 所在乡镇/街道（此街道为社区街道，不是道路信息）
	Towncode         string `json:"towncode"`          // 乡镇街道编码
	SeaArea          string `json:"seaArea"`           // 所属海域信息（可能没有）
	FormattedAddress string `json:"formatted_address"` // 格式化后的详细地址（可直接展示）
}

// ///////////////////////////////////////////////////// 逆地理编码结构end

func (c *client) regeoBase(ctx context.Context, location *Location) (*regeoBaseResponse, error) {
	url := fmt.Sprintf(regeoExtensionsBaseUrl, conf.Get(constants.AMapAppKey), location.LocationStr)
	httpClient := http2.NewHttpClient()
	httpClient.SetHeader("Accept", "application/json, text/plain, */*")
	resp, err := httpClient.Get(ctx, url)
	if err != nil {
		trlogger.Errorf(ctx, "regeoBase http get err: [%+v]", err)
		return nil, err
	}

	var data regeoBaseResponse
	err = http2.ResToObj(resp, &data)
	if err != nil {
		trlogger.Errorf(ctx, "regeoBase http2.ResToObj err: [%+v]", err)
		return nil, err
	}

	if data.RegeoCode == nil {
		err = errors.New("regeoBase data.RegeoCode is nil")
		trlogger.Errorf(ctx, "%+v", err)
		return nil, err
	}

	if data.RegeoCode.AddressComponent == nil {
		err = errors.New("regeoBase data.RegeoCode.AddressComponent is nil")
		trlogger.Errorf(ctx, "%+v", err)
		return nil, err
	}

	return &data, nil
}

func (c *client) getRoutePlan(ctx context.Context, opt *GetRoutePlanOption) ([]string, error) {
	//wayPoint, err := utils.ToJsonString(opt.Waypoints)
	//if err != nil {
	//	trlogger.Errorf(ctx, "getRoutePlan ToJsonString err: [%+v]", err)
	//	return []string{}, err
	//}

	wayPoint := strings.Join(opt.Waypoints, ";")
	url := fmt.Sprintf(routePlanV1, conf.Get(constants.AMapAppKey), opt.Origin, opt.Destination, wayPoint)
	httpClient := http2.NewHttpClient()
	httpClient.SetHeader("Accept", "application/json, text/plain, */*")
	resp, err := httpClient.Get(ctx, url)
	if err != nil {
		trlogger.Errorf(ctx, "getRoutePlan http get err: [%+v]", err)
		return []string{}, err
	}
	fmt.Println(resp)
	var data getRoutePlanResponse
	err = http2.ResToObj(resp, &data)
	if err != nil {
		trlogger.Errorf(ctx, "getRoutePlan http2.ResToObj err: [%+v]", err)
		return []string{}, err
	}

	if data.Route == nil {
		trlogger.Errorf(ctx, "getRoutePlan data.Route is nil")
		return []string{}, nil
	}

	if len(data.Route.Paths) <= 0 {
		trlogger.Errorf(ctx, "getRoutePlan data.Route.Paths len <= 0")
		return []string{}, nil
	}
	paths := data.Route.Paths[0]
	var polylines []string
	for _, st := range paths.Steps {
		stepPolyline := strings.Split(st.Polyline, ";")
		polylines = append(polylines, stepPolyline...)
	}
	return polylines, nil
}

func (c *client) getLiveWeather(ctx context.Context, adCode int64) (*weatherLiveBase, error) {
	url := fmt.Sprintf(weatherLive, conf.Get(constants.AMapAppKey), adCode)
	httpClient := http2.NewHttpClient()
	httpClient.SetHeader("Accept", "application/json, text/plain, */*")
	resp, err := httpClient.Get(ctx, url)
	if err != nil {
		trlogger.Errorf(ctx, "getLiveWeather http get err: [%+v]", err)
		return nil, err
	}

	var data weatherLiveResponse
	err = http2.ResToObj(resp, &data)
	if err != nil {
		trlogger.Errorf(ctx, "getLiveWeather http2.ResToObj err: [%+v]", err)
		return nil, err
	}

	if len(data.Lives) <= 0 {
		trlogger.Errorf(ctx, "getLiveWeather data.Lives len < 0")
		return nil, err
	}

	return data.Lives[0], nil
}
