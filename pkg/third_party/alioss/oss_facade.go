package alioss

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/woaijssss/tros"
	"github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/constants"
	trlogger "github.com/woaijssss/tros/logx"
	"time"
)

var Client = new(client)

func (c *client) Init(atx tros.AppContext) error {
	c.BucketName = conf.GetString(constants.AliOssBucket)
	c.Endpoint = conf.GetString(constants.AliOssUrl)
	c.AccessKeyId = conf.GetString(constants.AliOssAccessKeyId)
	c.AccessKeySecret = conf.GetString(constants.AliOssAccessKeySecret)

	return nil
}

func (c *client) GenUploadUrl(ctx context.Context, key string) (*OssSignature, error) {
	ossSignature, err := c.generateOssSignature(c.BucketName, key, constants.GTimeout)
	if err != nil {
		trlogger.Errorf(ctx, "GenUploadUrl generateOssSignature err: [%+v]", err)
		return nil, err
	}

	return ossSignature, nil
}

// GetRemoteSignUrl
/*
	key: 从 bucket 名字开始的文件路径
*/
func (c *client) GetRemoteSignUrl(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", nil
	}
	bucket, err := c.newClientWithBucket(ctx, c.BucketName)
	if err != nil {
		trlogger.Errorf(ctx, "GetRemoteSignUrl newClientWithBucket newClient err: [%+v]", err)
		return "", err
	}

	url, err := bucket.SignURL(key, oss.HTTPGet, int64(time.Hour*24)) // 一天有效期
	if err != nil {
		trlogger.Errorf(ctx, "GetRemoteSignUrl SignURL err: [%+v]", err)
		return "", err
	}
	return url, nil
}

func (c *client) GetRemoteSignUrlWithNoError(ctx context.Context, path string) string {
	if path == "" {
		return ""
	}
	bucket, err := c.newClientWithBucket(ctx, c.BucketName)
	if err != nil {
		trlogger.Errorf(ctx, "GetRemoteSignUrlWithNoError newClientWithBucket newClient err: [%+v]", err)
		return ""
	}

	url, err := bucket.SignURL(path, oss.HTTPGet, int64(time.Hour*24)) // 一天有效期
	if err != nil {
		trlogger.Errorf(ctx, "GetRemoteSignUrlWithNoError SignURL err: [%+v]", err)
		return ""
	}
	return url
}
