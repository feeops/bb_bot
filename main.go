package main

import (
	"bb_bot/database"
	"context"
	"fmt"
	"github.com/imroc/req/v3"
)

var (
	bitBaseURL = "http://127.0.0.1:54345"

	client   = req.C()
	email    = "xjmctjytcxmoeg@outlook.com"
	password = "2%)2^w4Y$RyW9e8&lYnUrI"
)

func main() {
	var err error
	readConfig()
	db := database.InitDB()

	err = heathCheck()
	if err != nil {
		fmt.Printf("指纹浏览器健康检查请求失败:%s\n", err.Error())
		waitExit()
	}

	var IPUsed bool
	IP, usedTime, err := checkProxy()
	if err != nil {
		fmt.Printf("指纹浏览器健康检查请求失败:%s\n", err.Error())
		waitExit()
	}
	if len(usedTime) == 0 {
		fmt.Printf("ip:%s 未被使用过\n", IP)
		IPUsed = false
	} else {
		fmt.Printf("ip:%s 使用过,使用时间:%s \n", IP, usedTime)
		IPUsed = true
	}

	windowsName := "第一资本测试"
	var browserID string
	browserID = findBrowser(windowsName)
	if len(browserID) == 0 {
		fmt.Println("浏览器窗口ID获取失败")
		waitExit()
	}

	controlURL, err := openBrowser(browserID)
	if err != nil {
		fmt.Println("发送浏览器获取控制地址请求失败")
		waitExit()
	}
	if len(controlURL) == 0 {
		fmt.Println("发送浏览器获取控制地址失败")
		waitExit()
	}
	var refURL string
	balance, err := getBalance(controlURL, browserID)
	if err != nil {
		fmt.Printf("email: %s %s\n", email, err.Error())
		db.Account.Create().SetBalance(balance).SetWindowName(windowsName).
			SetEmail(email).SetPassword(password).SetIP(IP).SetRefURL(refURL).
			SetRemark(err.Error()).
			SetIPUsed(IPUsed).SaveX(context.Background())
		return
	}
	fmt.Printf("email: %s 余额: %.2f\n", email, balance)
	if mode == 1 {
		refURL = getRef(controlURL)
		fmt.Printf("email: %s 推广链接: %s\n", email, refURL)
	}

	db.Account.Create().SetBalance(balance).SetWindowName(windowsName).
		SetEmail(email).SetPassword(password).SetIP(IP).SetRefURL(refURL).
		SetIPUsed(IPUsed).SaveX(context.Background())

}
