package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisCLi struct{
	rdb *redis.Client
	rcdb *redis.ClusterClient
	redisType string
}

const(
	SingleRedis = "SingleRedis"
	ClusterRedis = "ClusterRedis"
)

func (redisCli *RedisCLi) initClient()(err error){
	redisCli.rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB:0,
	})

	_,err = redisCli.rdb.Ping().Result()
	if err!=nil{
		fmt.Println("connect failed")
		return 
	}
	fmt.Println("connect success")
	return 
}

func (redisCli *RedisCLi) initClusterClient()(err error){
	redisCli.rcdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"127.0.0.1:6379","127.0.0.1:6380","127.0.0.1:6381"},
		Password: "",
	})

	_,err = redisCli.rcdb.Ping().Result()
	if err!=nil{
		fmt.Println("connect cluster failed")
		return 
	}
	fmt.Println("connect cluster success")
	return 
}

func (redisCli *RedisCLi) set(key string,value string)bool{
	if SingleRedis == redisCli.redisType{
		return redisCli.setSingleRedis(key,value)
	}else{
		return redisCli.setClusterRedis(key,value)
	}
}

func (redisCli *RedisCLi) setSingleRedis(key string,value string)bool{
	redisCli.initClient()
	redisCli.rdb.Set(key,value,0)
	return true
}
func (redisCli *RedisCLi) setClusterRedis(key string,value string)bool{
	redisCli.initClusterClient()
	redisCli.rcdb.Set(key,value,0)
	return true
}

func (redisCli *RedisCLi) get(key string)string{
	if SingleRedis == redisCli.redisType{
		return redisCli.getSingleRedis(key)
	}else{
		return redisCli.getClusterRedis(key)
	}
}

func (redisCli *RedisCLi) getSingleRedis(key string)string{
	redisCli.initClient()
	value,err:=redisCli.rdb.Get(key).Result()
	if err==redis.Nil{
		fmt.Println("key not exist")
		return ""
	}else if err !=nil{
		fmt.Println("get key failed err:",err,";value=",value)
		return ""
	}
	return value
}
func (redisCli *RedisCLi) getClusterRedis(key string)string{
	redisCli.initClusterClient()
	value,err:=redisCli.rcdb.Get(key).Result()
	if err==redis.Nil{
		fmt.Println("key not exist")
		return ""
	}else if err !=nil{
		fmt.Println("get key failed err:",err,";value=",value)
		return ""
	}
	return value
}