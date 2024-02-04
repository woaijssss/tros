package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
)

type TokenInfo struct {
	Uuid     string `json:"uuid"`
	UserName string `json:"user_name"`
	UserId   string `json:"user_id"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password"`
	Expire   string `json:"expire"`
}

func CreateToken(key string, tokenInfo TokenInfo) string {

	m := make(map[string]interface{}, 0)
	m["uuid"] = tokenInfo.Uuid
	m["user_name"] = tokenInfo.UserName
	m["user_id"] = tokenInfo.UserId
	m["role"] = tokenInfo.Role
	m["password"] = tokenInfo.Password
	m["expire"] = tokenInfo.Expire
	m["phone"] = tokenInfo.Phone

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}

	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func ParseToken(key string, tokenString string) (interface{}, bool) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if token == nil || token.Valid == false || token.Claims == nil {
		return "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return "", false
	}
}

func GetOriginalsInfo(claims interface{}) *TokenInfo {

	tokenInfo := TokenInfo{}
	v := reflect.ValueOf(claims)
	if v.Kind() != reflect.Map {
		return &tokenInfo
	}

	uuid, ok := claims.(jwt.MapClaims)["uuid"].(string)
	if ok {
		tokenInfo.Uuid = uuid
	}

	userName, ok := claims.(jwt.MapClaims)["user_name"].(string)
	if ok {
		tokenInfo.UserName = userName
	}

	userId, ok := claims.(jwt.MapClaims)["user_id"].(string)
	if ok {
		tokenInfo.UserId = userId
	}

	role, ok := claims.(jwt.MapClaims)["role"].(string)
	if ok {
		tokenInfo.Role = role
	}

	password, ok := claims.(jwt.MapClaims)["password"].(string)
	if ok {
		tokenInfo.Password = password
	}

	expire, ok := claims.(jwt.MapClaims)["expire"].(string)
	if ok {
		tokenInfo.Expire = expire
	}

	phone, ok := claims.(jwt.MapClaims)["phone"].(string)
	if ok {
		tokenInfo.Phone = phone
	}

	return &tokenInfo
}
