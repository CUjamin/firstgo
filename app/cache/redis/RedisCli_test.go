package redis

import (
	"fmt"
	"testing"
)


func TestSet(t *testing.T){
	var redisCli RedisCLi = RedisCLi{redisType: SingleRedis}
	var result = redisCli.set("S123","S123")
	fmt.Println("Set success=",result)
}

func TestGet(t *testing.T){
	var redisCli RedisCLi = RedisCLi{redisType: SingleRedis}
	var value = redisCli.get("S123")
	fmt.Println("value=",value)
}

func TestCSet(t *testing.T){
	var redisCli RedisCLi = RedisCLi{redisType: ClusterRedis}
	var result = redisCli.set("C123","C123")
	fmt.Println("Set success=",result)
}

func TestCGet(t *testing.T){
	var redisCli RedisCLi = RedisCLi{redisType: ClusterRedis}
	var value = redisCli.get("C123")
	fmt.Println("value=",value)
}
