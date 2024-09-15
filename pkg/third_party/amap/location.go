package amap

import (
	"context"
	"errors"
	trlogger "gitee.com/idigpower/tros/logx"
	"gitee.com/idigpower/tros/pkg/utils"
	"strconv"
	"strings"
)

const locationLen = 2 // 坐标信息长度，必须有经纬度，形如 116.405668,39.898906

func parseLocation(ctx context.Context, locationStr string) (*Location, error) {
	var err error
	if locationStr == "" {
		err = errors.New("locationStr is empty")
		trlogger.Errorf(ctx, "ParseLocation err: [%+v]", err)
		return nil, err
	}
	l := strings.Split(locationStr, ",")
	if len(l) != locationLen {
		err = errors.New("locationStr split len is not 2")
		trlogger.Errorf(ctx, "ParseLocation err: [%+v]", err)
		return nil, err
	}
	longitude, err := strconv.ParseFloat(l[0], 64)
	if err != nil {
		return nil, err
	}
	latitude, err := strconv.ParseFloat(l[1], 64)
	if err != nil {
		return nil, err
	}

	return &Location{
		LocationStr: locationStr,
		Longitude:   utils.FloatRetain(longitude, 6),
		Latitude:    utils.FloatRetain(latitude, 6),
	}, nil
}
