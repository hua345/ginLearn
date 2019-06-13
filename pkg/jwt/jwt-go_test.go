package jwt

import (
	"fmt"
	"testing"
)

// 单元测试
// go test -v
func TestCreateToken(t *testing.T) {
	helloMap := map[string]string{
		"name":   "fang",
		"userId": "001",
	}
	tokenStr := CreateToken(helloMap)
	fmt.Println("Token: " + tokenStr)
	tokenStrWithKey := CreateToken(helloMap, "key1", "key2")
	fmt.Println("TokenStrWithKey: " + tokenStrWithKey)
}

func TestParseToken(t *testing.T) {
	helloMap := map[string]string{
		"name":   "fang",
		"userId": "001",
	}
	tokenStr := CreateToken(helloMap)
	fmt.Println("Token: " + tokenStr)
	resultMap, ok := ParseToken(tokenStr)
	if ok {
		for key, value := range resultMap {
			fmt.Println("key:" + key + " value: " + value)
		}
	}
}
