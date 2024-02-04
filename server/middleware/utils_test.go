package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestMarkRequestFromGRpcGateway(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	MarkRequestFromGRpcGateway(req)
	assert.Equal(t, "true", req.Header.Get(upperGRpcGatewayFlag))
}

func TestIsRequestFromGRpcGateway(t *testing.T) {
	md := metadata.Pairs()
	assert.False(t, IsRequestFromGRpcGateway(metautils.NiceMD(md)))
	md.Set(upperGRpcGatewayFlag, "true")
	assert.True(t, IsRequestFromGRpcGateway(metautils.NiceMD(md)))
}
