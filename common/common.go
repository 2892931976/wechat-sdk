package common

const (
	UnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder" //微信统一下单
)

const (
	AccessTokenUrl  = "https://api.weixin.qq.com/sns/oauth2/access_token"  //code获取access_token
	RefreshTokenUrl = "https://api.weixin.qq.com/sns/oauth2/refresh_token" //重新获取access_token
	UserInfoUrl     = "https://api.weixin.qq.com/sns/userinfo"             //通过access_token获取userInfo
)
