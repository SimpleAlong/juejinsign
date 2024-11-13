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

type TodayStatus struct {
	ErrMsg string `json:"err_msg"`
	ErrNo  int    `json:"err_no"`
	Data   TodayStatusData
}
type TodayStatusData struct {
	CheckInDone bool `json:"check_in_done"`
	LtTaskExist bool `json:"lt_task_exist"`
}

var preUrl = "https://api.juejin.cn"

type Sign struct {
	aid     int64
	uuid    int64
	cookie  string
	msToken string
	bogus   string
	token   string
}

func New() *Sign {
	return new(Sign)
}
func (mod *Sign) AddAid(aid int64) {
	mod.aid = aid
}
func (mod *Sign) AddUuid(uuid int64) {
	mod.uuid = uuid
}
func (mod *Sign) AddCookie(cookie string) {
	mod.cookie = cookie
}
func (mod *Sign) AddMsToken(msToken string) {
	mod.msToken = msToken
}

func (mod *Sign) AddBogus(bogus string) {
	mod.bogus = bogus
}
func (mod *Sign) AddToken(token string) {
	mod.token = token
}
func (mod *Sign) Do() {
	isSign := mod.checkSign()
	mod.sign(isSign)
	mod.Lottery()
}
func (mod *Sign) sign(isSign bool) {
	if !isSign {
		log.Println("今日已签到")
		return
	}
	req := &http.Request{Header: mod.header(), Method: "POST", Body: io.NopCloser(strings.NewReader("{}"))}

	req.URL, _ = url.Parse(preUrl + fmt.Sprintf("/growth_api/v1/check_in?aid=%d&uuid=%d&spider=0&msToken=%s&a_bogus=%s", mod.aid, mod.uuid, mod.msToken, mod.bogus))
	var client = http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	var sign struct {
		ErrMsg string `json:"err_msg"`
		ErrNo  int    `json:"err_no"`
	}
	content, _ := io.ReadAll(rsp.Body)
	err = json.Unmarshal(content, &sign)
	log.Println("今日签到", err)
}

func (mod *Sign) checkSign() bool {
	req := &http.Request{Header: mod.header(), Method: "GET"}
	req.URL, _ = url.Parse(preUrl + fmt.Sprintf("/growth_api/v2/get_today_status?aid=%d&uuid=%d&spider=0", mod.aid, mod.uuid))
	var client = http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return false
	}
	data := new(TodayStatus)
	content, _ := io.ReadAll(rsp.Body)
	err = json.Unmarshal(content, data)
	if err != nil {
		return false
	}
	if data.ErrNo != 0 || data.Data.CheckInDone {
		return false
	}
	return true
}
