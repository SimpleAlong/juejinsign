package juejinsign

import (
	"net/http"
	"net/url"
)

func (mod *Sign) request(method string, url2 *url.URL) (*http.Response, error) {
	req := &http.Request{Header: mod.header(), Method: method}
	req.URL = url2
	var client = http.Client{}
	return client.Do(req)
}
func (mod *Sign) header() (h http.Header) {
	h = make(http.Header)

	mod.handleHeader(h)
	return
}
func (mod *Sign) handleHeader(h http.Header) {
	h.Set("accept", "*/*")
	h.Set("accept-language", "zh-CN,zh;q=0.9")
	h.Set("content-type", "application/json")
	h.Set("cookie", mod.cookie)
	h.Set("origin", "https://juejin.cn")
	h.Set("priority", "u=1, i")
	h.Set("referer", "https://juejin.cn/")
	h.Set("sec-ch-ua", "\"Chromium\";v=\"130\", \"Google Chrome\";v=\"130\", \"Not?A_Brand\";v=\"99\"")
	h.Set("sec-ch-ua-mobile", "?0")
	h.Set("sec-ch-ua-platform", "macOS")
	h.Set("sec-fetch-dest", "empty")
	h.Set("sec-fetch-mode", "cors")
	h.Set("sec-fetch-site", "same-site")
	h.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	h.Set("x-secsdk-csrf-token", mod.token)
}
