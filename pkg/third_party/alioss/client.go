package alioss

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
	"time"
)

type client struct {
	BucketName      string // 存储桶名称
	Endpoint        string // oss访问域名
	AccessKeyId     string // oss访问key id
	AccessKeySecret string // oss访问key secret
}

func (c *client) newClient(ctx context.Context) (*oss.Client, error) {
	cli, err := oss.New(c.Endpoint, c.AccessKeyId, c.AccessKeySecret)
	if err != nil {
		trlogger.Errorf(ctx, "alioss newClient err: [%+v]", err)
		return nil, err
	}

	return cli, nil
}

func (c *client) newClientWithBucket(ctx context.Context, bucketName string) (*oss.Bucket, error) {
	cli, err := oss.New(c.Endpoint, c.AccessKeyId, c.AccessKeySecret)
	if err != nil {
		trlogger.Errorf(ctx, "alioss newClient err: [%+v]", err)
		return nil, err
	}

	return c.getBucket(ctx, cli, bucketName)
}

type OssSignature struct {
	SignUrl     string
	AccessKeyId string `json:"accessKeyId"`
	//Host           string `json:"host"`
	Policy    string `json:"policy"`
	Signature string `json:"signature"`
	//SecurityToken  string `json:"securityToken"`
	//ExpirationTime string `json:"expirationTime"`
}

func generateSignature(message string, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(message))
	signature := hash.Sum(nil)
	return encrypt.Base64Encode(string(signature))
}

func (c *client) generateOssSignature(host, objectName string, expirationTime time.Duration) (*OssSignature, error) {

	// 生成策略
	//policy := fmt.Sprintf(`{"Version":"1","Statement":[{"Action":["oss:PutObject"],"Effect":"Allow","Resource":"acs:oss:%s:%s:%s"}]}`, host, c.BucketName, objectName)
	//policy := fmt.Sprintf(`{"Version":"1","Statement":[{"Action":["oss:PutObject"],"Effect":"Allow","Resource":"acs:oss:*:*:%s"}]}`, host)
	//policy := fmt.Sprintf(`{"expiration": "2034-01-01T12:00:00.000Z","conditions":[["content-length-range", 0, 1048576000]]}]}`)
	//policyB64 := encrypt.Base64Encode(policy)
	// 生成签名URL
	policyText := map[string]interface{}{
		"expiration": getFullTimeoutZ(expirationTime),
		"conditions": [][]interface{}{
			{"content-length-range", 0, 1048576000},
		},
	}
	policyTextJson, _ := json.Marshal(policyText)
	policy := encrypt.Base64Encode(string((policyTextJson)))
	// 生成签名
	signature := generateSignature(policy, c.AccessKeySecret)

	return &OssSignature{
		AccessKeyId: c.AccessKeyId,
		Policy:      policy,
		Signature:   signature,
		SignUrl:     fmt.Sprintf("https://%s.oss-cn-beijing.aliyuncs.com", host),
	}, nil
}

func (c *client) getBucket(ctx context.Context, cli *oss.Client, bucketName string) (*oss.Bucket, error) {
	bucket, err := cli.Bucket(bucketName)
	if err != nil {
		trlogger.Errorf(ctx, "alioss getBucket err: [%+v]", err)
		return nil, err
	}
	return bucket, nil
}

func getFullTimeoutZ(expirationTime time.Duration) string {
	return time.Now().Add(expirationTime).Format("2006-01-02T15:04:05.000Z")
}
