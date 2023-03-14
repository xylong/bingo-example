package user

// ThirdParty 第三方信息
type ThirdParty struct {
	WechatUnionid        string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信unionid" json:"-" bson:"wechat_unionid"`
	WechatAppletOpenid   string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信小程序🆔" json:"-" bson:"wechat_applet_openid"`
	WechatOfficialOpenid string `gorm:"type:varchar(100);default:null;uniqueIndex;comment:微信公众号🆔" json:"-" bson:"wechat_official_openid"`
}

func NewThirdParty() *ThirdParty {
	return &ThirdParty{}
}
