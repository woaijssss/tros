package wechat

import (
	"fmt"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
)

func (c *client) Signature(s interface{}, skipKeys []string) string {
	signature := utils.StructToXMLKeyValueSorted(s, skipKeys)
	stringSignTemp := fmt.Sprintf("%s&key=%s", signature, c.wechatPayApiV2Key)
	return encrypt.EncodeMD5Upper(stringSignTemp)
}
