package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

func main() {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.56.11:7000", "192.168.56.11:7001", "192.168.56.11:7002", "192.168.56.11:7003", "192.168.56.11:7004", "192.168.56.11:7005"},
	})
	redisdb.Set("name", "laowang", time.Duration(4)*time.Hour)
	go func() {
		var x = 1;
		for {
			pr, err := redisdb.Ping().Result()
			if err == nil {
				fmt.Printf("ping redis cluster is : %s\n", pr)
				fmt.Println("redisCluster info:")
				rinfo, _ := redisdb.ClusterInfo().Result()
				fmt.Println(rinfo)
				rnode, _ := redisdb.ClusterNodes().Result()
				fmt.Println("redisNodes info:")
				fmt.Println(rnode)
			}
			name, err2 := redisdb.Get("name").Result()
			if err2 != nil {
				fmt.Errorf("name值不存在\n")
			}
			fmt.Printf("name=%s\n", name)
			xage := fmt.Sprintf("age%d", x)
			sr, err3 := redisdb.Set(xage, x, time.Duration(5)*time.Second).Result()
			if err3 != nil {
				fmt.Errorf("redis集群写入值失败: %s\n", err3)
			}
			fmt.Println("测试redis集群写入数据:")
			fmt.Printf("> set %s %d\n", xage, x)
			fmt.Println(sr)
			fmt.Println(">")
			gr, _ := redisdb.Get(xage).Result()
			fmt.Printf("> get %s\n", xage)
			fmt.Println(gr)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			time.Sleep(time.Duration(10) * time.Second)
			x++;
		}
	}()
	time.Sleep(time.Duration(1) * time.Hour)
}
