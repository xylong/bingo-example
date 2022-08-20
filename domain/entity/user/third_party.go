package user

// ThirdParty 第三方信息
type ThirdParty struct {
	WechatUnionid        string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信unionid" json:"wechat_unionid" bson:"wechat_unionid"`
	WechatAppletOpenid   string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信小程序🆔" json:"wechat_applet_openid" bson:"wechat_applet_openid"`
	WechatOfficialOpenid string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信公众号🆔" json:"wechat_official_openid" bson:"wechat_official_openid"`
}

func NewThirdParty() *ThirdParty {
	return &ThirdParty{}
}
