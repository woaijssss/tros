package utils

import (
	"fmt"
	"gitee.com/idigpower/tros/trerror"
	"github.com/golang-jwt/jwt"
	"reflect"
)

type TokenInfo struct {
	Uuid string `json:"uuid"`
	//UserName string `json:"user_name"`
	UserId int64 `json:"user_id"`
	//Phone    string `json:"phone"`
	//Role     string `json:"role"`
	//Password string `json:"password"`
	//Expire   string `json:"expire"`
}

func CreateToken(key string, tokenInfo *TokenInfo) string {
	m := make(map[string]interface{}, 0)
	//m["uuid"] = tokenInfo.Uuid
	//m["user_name"] = tokenInfo.UserName
	m["user_id"] = tokenInfo.UserId
	//m["role"] = tokenInfo.Role
	//m["password"] = tokenInfo.Password
	//m["expire"] = tokenInfo.Expire
	//m["phone"] = tokenInfo.Phone

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func ParseToken(key string, tokenString string) (*TokenInfo, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ParseToken Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if token == nil || token.Valid == false || token.Claims == nil {
		return nil, trerror.DefaultTrError("token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, trerror.DefaultTrError(fmt.Sprintf("ParseToken convert MapClaims fail: [%+v]", token.Claims))
	}
	tokenInfo, err := MapToJson[TokenInfo](claims)
	if err != nil {
		return nil, trerror.DefaultTrError(fmt.Sprintf("ParseToken token map to struct fail: [%+v]", err))
	}
	return &tokenInfo, nil
}

// 解析JWT，但不验证签名
func ParseTokenWithoutVerify(tokenString string) (*TokenInfo, error) {
	// golang和java服务通用的userId key
	token, err := jwt.Parse(tokenString, nil)
	if err != nil {
		if _, ok := err.(*jwt.ValidationError); !ok {
			return nil, trerror.DefaultTrError(fmt.Sprintf("Error parsing token: [%+v]", err))
		}
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, trerror.DefaultTrError(fmt.Sprintf("ParseTokenWithoutVerify convert MapClaims fail: [%+v]", token.Claims))
	}
	tokenInfo, err := MapToJson[TokenInfo](claims)
	if err != nil {
		return nil, trerror.DefaultTrError(fmt.Sprintf("ParseTokenWithoutVerify token map to struct fail: [%+v]", err))
	}
	return &tokenInfo, nil
}

func GetOriginalsInfo(claims interface{}) *TokenInfo {

	tokenInfo := TokenInfo{}
	v := reflect.ValueOf(claims)
	if v.Kind() != reflect.Map {
		return &tokenInfo
	}

	//uuid, ok := claims.(jwt.MapClaims)["uuid"].(string)
	//if ok {
	//	tokenInfo.Uuid = uuid
	//}

	//userName, ok := claims.(jwt.MapClaims)["user_name"].(string)
	//if ok {
	//	tokenInfo.UserName = userName
	//}

	userId, ok := claims.(jwt.MapClaims)["user_id"].(int64)
	if ok {
		tokenInfo.UserId = userId
	}

	//role, ok := claims.(jwt.MapClaims)["role"].(string)
	//if ok {
	//	tokenInfo.Role = role
	//}

	//password, ok := claims.(jwt.MapClaims)["password"].(string)
	//if ok {
	//	tokenInfo.Password = password
	//}

	//expire, ok := claims.(jwt.MapClaims)["expire"].(string)
	//if ok {
	//	tokenInfo.Expire = expire
	//}

	//phone, ok := claims.(jwt.MapClaims)["phone"].(string)
	//if ok {
	//	tokenInfo.Phone = phone
	//}

	return &tokenInfo
}
