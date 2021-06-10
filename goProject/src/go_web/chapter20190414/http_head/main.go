package main

import (
	"fmt"
	"net/http"
	"net"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {

	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport {
				Dial:func(network, addr string) (net.Conn, error){
					timeout := time.Second*2
					//http超时时间为2s超时
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		/**
		1 http.Head(url)是http包的一个方法
		
		2 func (c *Client) Head(url string) (resp *Response, err error)
		是绑定到Client这个http客户端-结构体实例的一个方法

		Client结构体介绍：
		1 Client类型代表HTTP客户端。它的零值（DefaultClient）是一个
		可用的使用DefaultTransport的客户端。

		2 Client的Transport字段一般会含有内部状态（缓存TCP连接），
		因此Client类型值应尽量被重用而不是每次需要都创建新的。
		Client类型值可以安全的被多个go程同时使用。

		3 Client类型的层次比RoundTripper接口（如Transport）高，
		还会管理HTTP的cookie和重定向等细节。
		*/
		// Head向指定的URL发出一个HEAD请求，返回Response代表一个HTTP请求的回复
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("head succ, status:%v\n", resp.Status)
	}
}
