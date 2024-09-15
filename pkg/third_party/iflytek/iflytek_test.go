package iflytek

import (
	"context"
	"fmt"
	"gitee.com/idigpower/tros"
	trlogger "gitee.com/idigpower/tros/logx"
	"os"
	"testing"
)

func TestStt(t *testing.T) {
	tros.New(
		tros.WithInitializers([]tros.Initializer{
			Client,
		}...),
	)

	filePath := "./test.wav"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfo, _ := file.Stat()

	ctx := context.Background()
	uploadResp, err := Client.Upload(ctx, &UploadReq{
		FilePath: file.Name(),
		FileSize: fileInfo.Size(),
		Duration: 200,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("orderId: ", uploadResp.OrderId)
	fmt.Println("failedReason: ", uploadResp.FailedReason)

	txt, err := Client.GetResult(ctx, uploadResp.OrderId)
	trlogger.Infof(context.Background(), "txt: %s", txt)
	trlogger.Infof(context.Background(), "err: %+v", err)
}
