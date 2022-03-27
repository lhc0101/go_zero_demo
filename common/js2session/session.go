package js2session

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

// UserSession 用户会话信息
//type UserSession struct {
//	OpenID     string `json:"openid"`      //用户唯一标识
//	SessionKey string `json:"session_key"` //会话密钥
//	UnionID    string `json:"unionid"`     //用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
//}
type UserSession struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

const WX_APP_SECRET = "4f6c85c0c6eca1f1016d5bb45f3c16a5"

// Code2session 根据免登授权码获取用户信息
func Code2session(code, appId, appSecret string) (*UserSession, error) {
	client := &http.Client{}

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code"
	r := strings.NewReplacer(
		"APPID", appId,
		"SECRET", appSecret,
		"JSCODE", code,
	)

	var result UserSession
	u := r.Replace(url)
	req, err := http.NewRequest("GET", u, nil)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err = jsoniter.Unmarshal([]byte(body), &result); err != nil {
		println("request parse", "url", u, "err", err)
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}
