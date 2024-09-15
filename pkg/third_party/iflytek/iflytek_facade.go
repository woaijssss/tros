package iflytek

import (
	"context"
	"gitee.com/idigpower/tros"
	"gitee.com/idigpower/tros/conf"
	"gitee.com/idigpower/tros/constants"
	trlogger "gitee.com/idigpower/tros/logx"
	"os"
)

var Client = new(client)

func (c *client) Init(atx tros.AppContext) error {
	Client.appId = conf.GetString(constants.IFlyTekId)
	Client.secretKey = conf.GetString(constants.IFlyTekSecretKey)

	return nil
}

type UploadReq struct {
	FilePath string
	Duration int64
	FileSize int64
}

type UploadResp struct {
	OrderId      string
	FailedReason string
}

type basicIflytek struct {
	Type      string
	AppId     string
	SecretKey string
	Host      string
	//AccessKey       string
	//AccessKeySecret string
	//Enable bool
	//CallbackUrl     string
}

const (
	lfasrHost    = "https://raasr.xfyun.cn/v2/api"
	apiUpload    = "/upload"
	apiGetResult = "/getResult"
)

func (c *client) Upload(ctx context.Context, req *UploadReq) (*UploadResp, error) {
	file, err := os.Open(req.FilePath)
	if err != nil {
		trlogger.Errorf(ctx, "sdk OpenFile err: %v", err)
		return nil, err
	}

	defer func() {
		deferErr := file.Close()
		if deferErr != nil {
			trlogger.Errorf(ctx, "Upload defer file.Close err: %v", deferErr)
		}
	}()
	orderId, failedReason, err := c.upload(ctx, &basicUpload{
		uploadUrl: lfasrHost + apiUpload,
		filePath:  req.FilePath,
		fileSize:  req.FileSize,
		duration:  req.Duration,
		f:         file,
	})
	if err != nil {
		return nil, err
	}

	return &UploadResp{
		OrderId:      orderId,
		FailedReason: failedReason,
	}, nil
}

// GetResult 获取科大讯飞的识别结果 api结果内容,音频识别内容,失败原因,error
func (c *client) GetResult(ctx context.Context, orderId string) (string, error) {
	return c.getResult(ctx, &basicGetResult{
		getResultUrl: lfasrHost + apiGetResult,
		orderId:      orderId,
	})
}
