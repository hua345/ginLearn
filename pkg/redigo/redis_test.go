package redigo

import (
	"ginLearn/pkg/config"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	config.InitRedisTest()
	Init()
	testKey := "name"
	testValue := "fang"
	SetStrWithExpire(testKey, testValue, 60*60)
	value, err := GetStr(testKey)
	if err != nil {
		panic(err)
	}
	if value != testValue {
		t.Log("Redis Get " + testKey + ": " + value)
		t.Error("Redis Get " + testKey + " != " + testValue)
	}
}
func BenchmarkSetStr(b *testing.B) {
	testKey := "name"
	testValue := "fang"
	b.StopTimer() //停止压力测试的时间计数
	config.InitRedisTest()
	Init()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		SetStrWithExpire(testKey, testValue, 60*60)
	}
}
func BenchmarkSet(b *testing.B) {
	testKey := "name"
	testValue := "fang"
	b.StopTimer() //停止压力测试的时间计数
	config.InitRedisTest()
	Init()
	conn := RedisPools.Get()
	defer conn.Close()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		_, err := conn.Do("SET", testKey, testValue, "EX", 60*60)
		if err != nil {
			panic(err)
		}
	}
}
func BenchmarkGetStr(b *testing.B) {
	testKey := "name"
	b.StopTimer() //停止压力测试的时间计数
	config.InitRedisTest()
	Init()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		GetStr(testKey)
	}
}
func BenchmarkGet(b *testing.B) {
	testKey := "name"
	b.StopTimer() //停止压力测试的时间计数
	config.InitRedisTest()
	Init()
	conn := RedisPools.Get()
	defer conn.Close()
	b.StartTimer()             //重新开始时间
	for i := 0; i < b.N; i++ { //use b.N for looping
		_, err := redis.String(conn.Do("GET", testKey))
		if err != nil {
			panic(err)
		}
	}
}
