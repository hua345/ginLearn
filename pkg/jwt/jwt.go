package jwt

import (
	"fmt"
	"ginLearn/internal/pkg"
	"ginLearn/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		util.EncodeMD5(username),
		util.EncodeMD5(password),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(pkg.AppConfig.JwtSecret)

	return token, err
}

// ParseMapToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return pkg.AppConfig.JwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// 创建token
func CreateMapToken(m map[string]string, keys ...string) string {
	key := pkg.AppConfig.JwtSecret
	if len(keys) > 0 {
		key = keys[0]
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

// 解析token
func ParseMapToken(tokenString string, keys ...string) (map[string]string, bool) {
	key := pkg.AppConfig.JwtSecret
	if len(keys) > 0 {
		key = keys[0]
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapData := make(map[string]string)
		for index, val := range claims {
			mapData[index] = fmt.Sprintf("%v", val)
		}
		return mapData, true
	} else {
		return nil, false
	}
}
