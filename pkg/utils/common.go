// Package utils add some context function
package utils

// GetLogKv will add some message to log
//func GetLogKv(c *gin.Context) []interface{} {
//	const capacity = 14
//
//	kv := make([]interface{}, 0, capacity)
//	kv = append(kv,
//		"uri", c.Request.URL.Path,
//		"query", c.Request.URL.RawQuery,
//		"method", c.Request.Method,
//		"status", c.Writer.Status(),
//		"req_id", c.GetHeader(middleware.HeaderRequestID),
//		"ip", c.ClientIP(),
//	)
//
//	t, uuid := GetIdentity(c)
//	if t != "" {
//		kv = append(kv, t, uuid)
//	}
//	return kv
//}

const (
	uuidCommonChar    = "0123456789"
	uuidFullUpperChar = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	uuidFullLowerChar = "0123456789abcdefghijklmnopqrstuvwxyz"
)
