package third_party

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
	"time"
)

func UploadPictureToCos(picturePath string) {
	secretID := ""
	secretKey := ""
	domain := "https://cloud.com"
	u, _ := url.Parse(domain)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  false,
				RequestBody:    false,
				ResponseHeader: false,
				ResponseBody:   false,
			},
		},
	})
	timeNow := time.Now()
	suffix := "bucket"
	filePath := fmt.Sprintf("/%s/%d%d%d/%s", suffix, timeNow.Year(), timeNow.Month(), timeNow.Day(), picturePath)
	v, _, err := c.Object.Upload(context.Background(), filePath, picturePath, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
