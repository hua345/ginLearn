package jwt

import (
	"fmt"
	"testing"
)

// 单元测试
// go test -v
func TestCreateMapToken(t *testing.T) {
	helloMap := map[string]string{
		"name":   "fang",
		"userId": "001",
	}
	tokenStr := CreateMapToken(helloMap)
	fmt.Println("Token: " + tokenStr)
	tokenStrWithKey := CreateMapToken(helloMap, "key1", "key2")
	fmt.Println("TokenStrWithKey: " + tokenStrWithKey)
}

func TestParseMapToken(t *testing.T) {
	helloMap := map[string]string{
		"name":   "fang",
		"userId": "001",
	}
	tokenStr := CreateMapToken(helloMap)
	fmt.Println("Token: " + tokenStr)
	resultMap, ok := ParseMapToken(tokenStr)
	if ok {
		for key, value := range resultMap {
			fmt.Println("key:" + key + " value: " + value)
		}
	}
}
