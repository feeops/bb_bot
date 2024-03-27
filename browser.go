package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/utils"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
)

func getPort(browserID string) (int64, error) {
	var remotePort int64
	resp, err := client.R().Post(bitBaseURL + "/browser/ports")
	if err != nil {
		return remotePort, err
	}

	for k, v := range gjson.Get(resp.String(), "data").Map() {
		if k == browserID {
			remotePort = v.Int()
			break
		}
	}

	return remotePort, err
}

func login(controlURL string) {
	u := launcher.MustResolveURL(controlURL)

	browser := rod.New().NoDefaultDevice().ControlURL(u).MustConnect()

	page := browser.MustPage("https://capitaloneshopping.com/sign-in").MustWaitLoad()
	utils.Sleep(interval)

	page.MustElement(`#id_email`).MustSelectAllText().MustInput(email)
	utils.Sleep(interval)
	page.MustElement(`#id_password`).MustSelectAllText().MustInput(password)
	utils.Sleep(interval)
	page.MustElementR("button", "Sign In").MustClick()
	utils.Sleep(interval)

}

func findBrowser(name string) (browserID string) {
	data := map[string]interface{}{
		"page":     0,
		"pageSize": 100,
		"name":     name,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(bitBaseURL + "/browser/list")
	if err != nil {
		fmt.Printf("指纹浏览器浏览器列表请求失败:%s\n", err.Error())
		waitExit()
	}

	for _, item := range gjson.Get(resp.String(), "data.list").Array() {
		browserID = item.Get("id").Str
		break
	}

	return browserID

}

func checkProxy() (IP string, usedTime string, err error) {
	data := map[string]interface{}{
		"host":          "127.0.0.1",
		"port":          7890,
		"proxyType":     "socks5",
		"proxyUserName": "admin",
		"proxyPassword": "admin",
		"checkExists":   1,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(bitBaseURL + "/checkagent")
	if err != nil {
		return "", "", err
	}

	respStr := resp.String()
	IP = gjson.Get(respStr, "data.data.ip").Str
	usedTime = gjson.Get(respStr, "data.data.usedTime").Str

	return IP, usedTime, nil

}

func heathCheck() error {

	resp, err := client.R().Post(bitBaseURL + "/health")
	if err != nil {
		return err
	}

	if gjson.Get(resp.String(), "success").Bool() == true {
		return nil
	} else {
		return fmt.Errorf("指纹浏览器健康检查失败:%s\n", resp.String())
	}

}

func getCookies(browserID string) {
	data := map[string]string{
		"browserId": browserID,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(bitBaseURL + "/browser/cookies/get")
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("browser cookies error")
		return
	}

	logger.Info().Str("resp", resp.String()).Msg("record cookies")

}

func openBrowser(browserID string) (string, error) {
	data := map[string]string{
		"id": browserID,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(bitBaseURL + "/browser/open")
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("browser cookies error")
		return "", err
	}
	logger.Info().Str("resp", resp.String()).Msg("openBrowser")
	return gjson.Get(resp.String(), "data.http").Str, nil
}

func browserDetail(browserID string) (float64, bool) {
	data := map[string]string{
		"id": browserID,
	}

	resp, err := client.R().SetBodyJsonMarshal(data).Post(bitBaseURL + "/browser/detail")
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("browser detail error")
		return 0, false
	}

	respStr := resp.String()
	userAgent := gjson.Get(respStr, "data.browserFingerPrint.userAgent").Str
	client.SetUserAgent(userAgent)
	cookie := gjson.Get(respStr, "data.cookie").Str
	for _, item := range gjson.Parse(cookie).Array() {
		domain := item.Get("domain").Str
		if strings.Contains(domain, "capitaloneshopping.com") {
		} else {
			continue
		}

		value := item.Get("value").Str

		// 解决net/http: invalid byte '"' in Cookie.Value; dropping invalid bytes问题
		if strings.Contains(value, `"`) {
			continue
		}

		name := item.Get("name").Str

		client.SetCommonCookies(&http.Cookie{
			Domain: domain,
			Name:   name,
			Path:   item.Get("path").Str,
			Value:  value,
		})

	}

	// logger.Info().Str("resp", respStr).Msg("browser detail info")
	proxyType := gjson.Get(respStr, "data.proxyType").Str
	host := gjson.Get(respStr, "data.host").Str
	port := gjson.Get(respStr, "data.port").Int()
	proxyUserName := gjson.Get(respStr, "data.proxyUserName").Str
	proxyPassword := gjson.Get(respStr, "data.proxyPassword").Str

	var ProxyURL string
	switch {

	case len(proxyUserName) > 0:
		ProxyURL = fmt.Sprintf("%s://%s:%s@%s:%d",
			proxyType, proxyUserName, proxyPassword, host, port)

	case proxyType == "noproxy":

	case len(proxyUserName) == 0:
		ProxyURL = fmt.Sprintf("%s://%s:%d",
			proxyType, host, port)
	}

	logger.Info().Str("ProxyURL", ProxyURL).Msg("ProxyURL")
	resp, err = client.SetProxyURL(ProxyURL).R().
		Get("https://capitaloneshopping.com/api/v1/lifetime_savings_info")
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("browser detail error")
		fmt.Println(err.Error())
		return 0, false
	}
	amount := gjson.Get(resp.String(), "credits.amount")

	return cast.ToFloat64(amount.Int()) / 100.0, amount.Exists()
}

func checkLogin(controlURL string) error {
	u := launcher.MustResolveURL(controlURL)

	browser := rod.New().NoDefaultDevice().ControlURL(u).MustConnect()

	URL := "https://capitaloneshopping.com/my-rewards/lifetime-savings"
	page := browser.MustPage(URL).MustWaitLoad()
	utils.Sleep(interval)
	html := page.MustHTML()
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		return err
	}

	node := htmlquery.FindOne(doc, `//span[@class="credits"]`)
	if node != nil {

	} else {
		login(controlURL)
	}

	return nil

}

func getBalance(controlURL string, browserID string) float64 {
	err := checkLogin(controlURL)
	if err != nil {
		fmt.Println(err)
		logger.Info().Str("error", err.Error()).Msg("checkLogin error")
		waitExit()
	}
	balance, exist := browserDetail(browserID)
	logger.Info().Float64("balance", balance).
		Bool("exist", exist).Msg("getBalance")
	if exist {
		return balance
	}

	return 0

}

func getRef(controlURL string) string {
	u := launcher.MustResolveURL(controlURL)

	browser := rod.New().NoDefaultDevice().ControlURL(u).MustConnect()

	page := browser.MustPage(couponURL).MustWaitLoad()
	utils.Sleep(interval)

	page.MustElementX(`//div[@class="tertiary-btn-small button-style"]`).MustClick()
	page.MustWaitLoad()
	utils.Sleep(interval)
	pages := browser.MustPages()
	for _, p := range pages {
		return p.MustInfo().URL
	}

	return ""
}
