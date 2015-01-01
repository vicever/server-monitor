// httptest
package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"net/http"
	"time"
)

func httprequest() {
	url_adr := []string{
		"http://www.baidu.com",
		"https://www.bestpay.com.cn",
		"http://www.zhanggs.com",
		"http://www.taobao.com",
	}
	for _, url := range url_adr {
		client := &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(netw, addr, time.Second*10)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(time.Second * 10))
					return conn, nil
				},
				ResponseHeaderTimeout: time.Second * 10,
			},
		}

		reqest, _ := http.NewRequest("GET", url, nil)

		//reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		//reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
		//reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
		//reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
		//reqest.Header.Set("Cache-Control", "max-age=0")
		//reqest.Header.Set("Connection", "keep-alive")
		response, err := client.Do(reqest)
		if err != nil {
			response.Body.Close()
			fmt.Println(err)

		} else {
			fmt.Println("The result of  ", url, "is :", response.StatusCode)

		}
		defer response.Body.Close()
		/*
			if response.StatusCode == 200 {
				body, _ := ioutil.ReadAll(response.Body)
				bodystr := string(body)
				fmt.Println(bodystr)
			}*/
	}
}

func main() {
	fmt.Println("Hello World!")
	httprequest()
}
