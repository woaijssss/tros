package middleware

import (
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

// X-GW-Flag 声明大小写变量是为了规避 Header.Set和MD.Get 因为大小写转换产生额外的内存申请
const (
	upperGRpcGatewayFlag = "X-GW-Flag"
	lowerGRpcGatewayFlag = "x-gw-flag"
)

// IsRequestFromGRpcGateway check gRpc request is from gRpc-Gateway
func IsRequestFromGRpcGateway(md metautils.NiceMD) bool {
	if flag := md.Get(lowerGRpcGatewayFlag); flag == "true" {
		return true
	}
	return false
}

// MarkRequestFromGRpcGateway mark http request is from gRpc-Gateway
func MarkRequestFromGRpcGateway(req *http.Request) {
	req.Header.Set(upperGRpcGatewayFlag, "true")
}

func ExcludePaths(excludes []string) map[string]struct{} {
	var skip map[string]struct{}

	if length := len(excludes); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, exclude := range excludes {
			skip[exclude] = struct{}{}
		}
	}

	return skip
}
