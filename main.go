package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"time"
)

var (
	URL     = "http://127.0.0.1:54345"
	headers = map[string]string{"Content-Type": "application/json"}

	client     = req.C()
	browserID  string
	remotePort int64
	email      = "fdeejkxxbbazyl@hotmail.com"
	password   = "VTTrNKyswK$A"
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

func findBrowser(name string) {
	data := map[string]interface{}{
		"page":     0,
		"pageSize": 100,
		"name":     name,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(URL + "/browser/list")
	if err != nil {
		fmt.Printf("指纹浏览器浏览器列表请求失败:%s\n", err.Error())
		waitExit()
	}

	for _, item := range gjson.Get(resp.String(), "data.list").Array() {
		browserID = item.Get("id").Str
		break
	}

	if len(browserID) == 0 {
		fmt.Println("浏览器窗口ID获取失败")
	}

	fmt.Println(browserID)

}

func getPort() {
	resp, err := client.R().Post(URL + "/browser/ports")
	if err != nil {
		fmt.Printf("指纹浏览器浏览器列表请求失败:%s\n", err.Error())
		waitExit()
	}

	for k, v := range gjson.Get(resp.String(), "data").Map() {
		if k == browserID {
			remotePort = v.Int()
			break
		}
	}

	fmt.Println(remotePort)
}

func control() {
	u := launcher.MustResolveURL(cast.ToString(remotePort))

	browser := rod.New().NoDefaultDevice().ControlURL(u).MustConnect()
	page := browser.MustPage("https://capitaloneshopping.com/sign-in").MustWaitLoad()
	utils.Sleep(5)
	page.MustElement(`#id_email`).MustSelectAllText().MustInput(email)
	utils.Sleep(5)
	page.MustElement(`#id_password`).MustSelectAllText().MustInput(password)
	utils.Sleep(5)
	page.MustElementR("button", "Sign In").MustClick()
	utils.Sleep(5)
	page.MustNavigate("https://capitaloneshopping.com/my-rewards/lifetime-savings").MustWaitLoad()

	err := rod.Try(func() {
		credit := page.Timeout(10 * time.Second).MustElementX(`//span[@class="credits"]`).MustText()
		fmt.Println(credit)
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(browser.MustGetCookies())
}

func main() {

	client.SetCommonHeaders(headers)
	heathCheck()
	checkProxy()
	findBrowser("第一资本测试")
	getPort()
	control()

}
