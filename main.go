package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
)

var (
	URL     = "http://127.0.0.1:54345"
	headers = map[string]string{"Content-Type": "application/json"}

	client = req.C()
)

func heathCheck() {

	resp, err := client.R().Post(URL + "/health")
	if err != nil {
		fmt.Printf("指纹浏览器健康检查请求失败:%s\n", err.Error())
		waitExit()
	}

	if gjson.Get(resp.String(), "success").Bool() == true {
		fmt.Println("指纹浏览器健康检查成功")
	} else {
		fmt.Printf("指纹浏览器健康检查失败:%s\n", resp.String())
		waitExit()
	}

}

func checkProxy() {
	data := map[string]interface{}{
		"host":          "127.0.0.1",
		"port":          7890,
		"proxyType":     "socks5",
		"proxyUserName": "admin",
		"proxyPassword": "admin",
		"checkExists":   1,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(URL + "/checkagent")
	if err != nil {
		fmt.Printf("指纹浏览器健康检查请求失败:%s\n", err.Error())
		waitExit()
	}

	respStr := resp.String()
	ip := gjson.Get(respStr, "data.data.ip").Str
	usedTime := gjson.Get(respStr, "data.data.usedTime").Str
	if len(usedTime) == 0 {
		fmt.Printf("ip:%s 未被使用过\n", ip)
	} else {
		fmt.Printf("ip:%s 使用过,使用时间:%s \n", ip, usedTime)
	}

	fmt.Println(resp.String())

}

func main() {

	client.SetCommonHeaders(headers)
	heathCheck()
	checkProxy()

}
