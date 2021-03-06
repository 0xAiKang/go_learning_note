
package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func GetConn() redis.Conn  {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		log.Println("connect redis error: ", err)
		return nil;
	}

	return conn
}

func main()  {
	conn := GetConn()
	if conn != nil {
		fmt.Println("connect redis success")
	}

	// 使用Do 方法执行Redis 命令
	// value1, err1 := redis.Values(conn.Do("lrange", "test_list", 0 , -1 ))
	value1, err1 := redis.Values(conn.Do("ZRANGE", "award_info", 0, -1, "WITHSCORES"))
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(value1)

	// 使用pipeline 方法执行Redis 命令
	// 往队列中添加指令
	conn.Send("GET", "hello")
	// 一次性将指令发送到服务器
	conn.Flush()
	// 接收服务端返回的数据
	v, err := conn.Receive()
	if err != nil {
		fmt.Println("redis get key error: ", err)
	}
	fmt.Println(v)
	fmt.Printf("%T ", v)
	s := v.([]byte)
	fmt.Printf("%T ", s)
	fmt.Println(string(s))
	// fmt.Println(string(v.([]byte)))

	// 取得连接之后，一定要记得关闭
	defer conn.Close()
}