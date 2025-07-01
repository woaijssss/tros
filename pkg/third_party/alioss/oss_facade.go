package alioss

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/woaijssss/tros"
	"github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/constants"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/trerror"
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

// DownloadFile 下载oss上的文件，保存到本地
func (c *client) DownloadFile(ctx context.Context, remotePath string, localSavedPath string) error {
	if remotePath == "" || localSavedPath == "" {
		return trerror.TR_ILLEGAL_OPERATION
	}
	bucket, err := c.newClientWithBucket(ctx, c.BucketName)
	if err != nil {
		trlogger.Errorf(ctx, "DownloadFile newClientWithBucket newClient err: [%+v]", err)
		return err
	}

	// 下载文件到本地。如果指定的本地文件不存在，则会被创建。如果存在，则会被覆盖。
	// 将OSS上的object名称为"your-object-name"的文件下载到本地路径"your-local-filename"
	err = bucket.GetObjectToFile(remotePath, localSavedPath)
	if err != nil {
		trlogger.Errorf(ctx, "DownloadFile bucket.GetObjectToFile err: [%+v][%s, %s]", err, remotePath, localSavedPath)
		return err
	}
	return nil
}

// UploadFile 将本地文件上传至oss
func (c *client) UploadFile(ctx context.Context, localPath string, remotePath string) (string, error) {
	if localPath == "" || remotePath == "" {
		return "", trerror.TR_ILLEGAL_OPERATION
	}
	bucket, err := c.newClientWithBucket(ctx, c.BucketName)
	if err != nil {
		trlogger.Errorf(ctx, "UploadFile newClientWithBucket newClient err: [%+v]", err)
		return "", err
	}

	fileSize, err := utils.GetCommonFileSize(localPath)
	if err != nil {
		trlogger.Errorf(ctx, "UploadFile utils.GetCommonFileSize err: [%+v][%s]", err, localPath)
		return "", err
	}

	fmt.Println("fileSize: ", fileSize)

	err = bucket.UploadFile(remotePath, localPath, oss.MinPartSize)
	if err != nil {
		trlogger.Errorf(ctx, "UploadFile bucket.UploadFile err: [%+v]", err)
		return "", err
	}

	return remotePath, nil
}
