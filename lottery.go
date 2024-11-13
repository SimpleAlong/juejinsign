package juejinsign

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (mod *Sign) Lottery() {
	req := &http.Request{Header: mod.header(), Method: "GET", Body: io.NopCloser(strings.NewReader("{}"))}
	req.URL, _ = url.Parse(preUrl + fmt.Sprintf("/growth_api/v1/lottery_config/get?aid=%d&uuid=%d&spider=0&msToken=%s&a_bogus=%s", mod.aid, mod.uuid, mod.msToken, mod.bogus))
	var client = http.Client{}
	rsp, err := client.Do(req)
	var conf struct {
		ErrMsg string `json:"err_msg"`
		ErrNo  int    `json:"err_no"`
		Data   struct {
			FreeCount int `json:"free_count"`
		}
	}
	content, _ := io.ReadAll(rsp.Body)
	err = json.Unmarshal(content, &conf)
	if conf.Data.FreeCount == 0 {
		return
	}
	req = &http.Request{Header: mod.header(), Method: "POST", Body: io.NopCloser(strings.NewReader("{}"))}
	req.URL, _ = url.Parse(preUrl + fmt.Sprintf("/growth_api/v1/lottery/draw?aid=%d&uuid=%d&spider=0&msToken=%s&a_bogus=%s", mod.aid, mod.uuid, mod.msToken, mod.bogus))
	rsp, err = client.Do(req)
	content, _ = io.ReadAll(rsp.Body)

	log.Println("免费抽奖", string(content), err)

}
