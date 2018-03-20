package login

import (
	"time"
	"errors"
	"net/url"
	"encoding/json"
	"wechat-sdk/utils"
	"wechat-sdk/common"
	"fmt"
)

type (
	WeConfig struct {
		AppId  string `json:"appid"`
		Secret string `json:"secret"`
	}

	AccessToken struct {
		AccessToken  string `json:"access_token,omitempty"`
		ExpiresIn    uint   `json:"expires_in,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
		OpenId       string `json:"openid,omitempty"`
		Scope        string `json:"scope,omitempty"`
		ErrCode      uint   `json:"errcode,omitempty"`
		ErrMsg       string `json:"errmsg,omitempty"`
		ExpiredAt    time.Time
	}

	WeUserInfo struct {
		OpenID     string `json:"openid,omitempty"`     // 授权用户唯一标识
		NickName   string `json:"nickname,omitempty"`   // 普通用户昵称
		Sex        uint32 `json:"sex,omitempty"`        // 普通用户性别，1为男性，2为女性
		Province   string `json:"province,omitempty"`   // 普通用户个人资料填写的省份
		City       string `json:"city,omitempty"`       //普通用户个人资料填写的城市
		Country    string `json:"country,omitempty"`    //国家，如中国为CN
		HeadImgUrl string `json:"headimgurl,omitempty"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
		//Privilege  string `json:"privilege"`
		UnionId string `json:"unionid,omitempty"` // 普通用户的标识，对当前开发者帐号唯一
		ErrCode uint   `json:"errcode,omitempty"`
		ErrMsg  string `json:"errmsg,omitempty"`
	}
)

// 微信APP登录 直接登录获取用户信息
func (m *WeConfig) AppLogin() {

}

// 微信小程序登录 直接登录获取用户信息
func (m *WeConfig) WexLogin() {

}

// 通过code获取AccessToken
func (m *WeConfig) GetWxAccessToken(code string) (accessToken *AccessToken, err error) {

	if code == "" {
		return accessToken, errors.New("getWxAccessToken error: code is null")
	}

	params := url.Values{
		"code":       []string{code},
		"grant_type": []string{"authorization_code"},
	}

	t, err := utils.Struct2Map(m)

	if err != nil {
		return accessToken, err
	}

	for k, v := range t {
		params.Set(k, v)
	}

	if err != nil {
		return accessToken, err
	}

	body, err := utils.NewRequest("GET", common.AccessTokenUrl, []byte(params.Encode()))


	fmt.Println(string(body))

	if err != nil {
		return accessToken, err
	}
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return accessToken, err
	}

	if accessToken.ErrMsg != "" {
		return accessToken, errors.New(accessToken.ErrMsg)
	}

	return
}

// 获取用户资料
func (m *AccessToken) GetUserInfo() (weUserInfo *WeUserInfo, err error) {
	if m.AccessToken == "" {
		return nil, errors.New("getWxUserInfo error: accessToken is null")
	}

	if m.OpenId == "" {
		return nil, errors.New("getWxUserInfo error: openID is null")
	}

	body, err := utils.NewRequest("GET", common.UserInfoUrl, []byte("access_token="+m.AccessToken+"&openid="+m.OpenId))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &weUserInfo)
	if err != nil {
		return nil, err
	}

	if weUserInfo.OpenID == "" {
		return weUserInfo, errors.New(weUserInfo.ErrMsg)
	}

	return
}
