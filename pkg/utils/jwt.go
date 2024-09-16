package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/woaijssss/tros/trerror"
	"reflect"
	"time"
)

const userTokenKey = "trlink.com"

type TokenInfo struct {
	Sub               string `json:"sub"`
	UserId            string `json:"user_id"`
	CurrentTimeMillis string `json:"currentTimeMillis"`
	Exp               int64  `json:"exp"`
	Iat               int64  `json:"iat"`
	Jti               string `json:"jti"`
}

func CreateToken(userNo string) string {
	now := time.Now().Unix()
	currentTimeMillis := now * 1000
	secretKey := fmt.Sprintf("%s%s%d", userNo, userTokenKey, currentTimeMillis)
	m := make(map[string]interface{}, 0)
	m["sub"] = userTokenKey
	m["user_id"] = userNo
	m["currentTimeMillis"] = fmt.Sprintf("%d", now*1000)
	m["exp"] = now + 850000000
	m["iat"] = now
	m["jti"] = userNo

	token := jwt.New(jwt.SigningMethodHS512)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(secretKey))
	return tokenString
}

func ParseToken(tokenString string) (*TokenInfo, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ParseToken Unexpected signing method: %v", token.Header["alg"])
		}
		//key := fmt.Sprintf("%s%s%d", userNo, userTokenKey, currentTimeMillis)
		return []byte(userTokenKey), nil
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
	token, _ := jwt.Parse(tokenString, nil)
	//if err != nil {
	//	if _, ok := err.(*jwt.ValidationError); !ok {
	//		return nil, trerror.DefaultTrError(fmt.Sprintf("Error parsing token: [%+v]", err))
	//	}
	//}
	if token == nil {
		return nil, trerror.DefaultTrError(fmt.Sprintf("Error parsing token"))
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

	userId, ok := claims.(jwt.MapClaims)["user_id"].(string)
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
