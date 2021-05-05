package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

/**
	ETCD 的哭护短
*/
func main() {
	// go 操作 ETCD 的步骤
	// 1. 通过clientv3.New 创建ETCD 客户端，链接ETCD
	// 2. 通过 步骤1 创建的 ETCD 客户端操作 ETCD
	// 3. 关闭ETCD链

	// 开启上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. 通过clientv3.New 创建 ETCD 客户端，链接 ETCD
	etcdConfig := clientv3.Config{
		// 链接终端，为一个字符串数组
		Endpoints: []string{"127.0.0.1:2379"},
		// 链接超过时间，5秒
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(etcdConfig)
	// 关闭链接
	defer cli.Close()

	if err != nil {
		panic("链接 ETCD 失败！原因为：" + err.Error())
	}

	// 设置值
	cli.Put(ctx, "/kkk/sex", "男")

	res,_ := cli.Get(ctx, "/kkk/name")

	fmt.Println("res.Kvs")
	fmt.Println(res.Kvs)


}