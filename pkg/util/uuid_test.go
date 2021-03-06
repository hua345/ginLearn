package util

import (
	"fmt"
	"testing"
)

// 单元测试
// go test -v
func TestGetUUID(t *testing.T) {
	fmt.Println("UUIDv4: " + GetUUID())
	uuidStr := GetUUID()
	fmt.Println("UUIDv4: " + uuidStr)
	if len(uuidStr) != 36 {
		t.Error(`UUID {uuidStr} length != 36`)
	}
}
