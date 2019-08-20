package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func createRedisConn () (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	return conn, err
}

var  redisPool = & redis.Pool{
	MaxIdle:30,
	MaxActive:1000,
	IdleTimeout: 5 * time.Second,
	Dial: func() (conn redis.Conn, err error) {
		return redis.Dial("tcp", "127.0.0.1:6379")
	},
}

func main() {
	//conn := redisPool.Get()

	//for i := 0; i < 1000000; i++ {
	//	//go run1(conn, i)
	//	go run2(i)
	//}

	run3()
}

func run3() {
	//连接redis
	conn, err := createRedisConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ok, err := conn.Do("SET", "name", "a")
	if err != nil {
		panic(err)
	}

	if ok != "OK" {
		panic("failed")
	}

}

func run1(conn redis.Conn, i int) {
	//通过Do函数，发送redis命令
	name := fmt.Sprintf("zhouweixin-%d", i)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	v, err := conn.Do("SET", "name", "zhouweixin-"+strconv.Itoa(i))
	if err != nil {
		panic(err)
	}

	if v != "OK" {
		return
	}

	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		panic(err)
	}

	if name != v {
		fmt.Printf("%s != %s\n", name, v)
	}
}

func run2(i int) {
	//连接redis
	conn, err := createRedisConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//通过Do函数，发送redis命令
	name := fmt.Sprintf("zhouweixin-%d", i)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	v, err := conn.Do("SET", "name", "zhouweixin-"+strconv.Itoa(i))
	if err != nil {
		panic(err)
	}

	if v != "OK" {
		return
	}

	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		panic(err)
	}

	if name != v {
		fmt.Printf("%s != %s\n", name, v)
	}
}