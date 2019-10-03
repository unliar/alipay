package alipay

import (
	"crypto/rsa"
	"fmt"
	ali "github.com/unliar/utils/go/alipay"
	"github.com/unliar/utils/go/http"
	"time"
)

type Client struct {
	AppID           string         // 应用id
	AliPayPublicKey *rsa.PublicKey // 支付宝公钥
	// 普通公钥签名可用
	AppPrivateKey *rsa.PrivateKey // 应用私钥
	// 公钥证书可以
	AppPublicKeyCert    string // app公钥证书
	AliPayPublicKeyCert string // 支付宝公钥证书
	AliPayRootCert      string // 支付宝根证书
	NotifyURL           string // 回调接口
	SignType            string // 签名类型
}

// 预下单接口
func (c *Client) TradePreCreate(p BizContentRequestParams) {
	v := Params{
		PublicRequestParams: PublicRequestParams{
			AppID:     c.AppID,
			Method:    AlipayTradePrecreateMethodName,
			Format:    DefaultFormat,
			Charset:   DefaultCharset,
			SignType:  c.SignType,
			NotifyURL: c.NotifyURL,
			Timestamp: time.Now().Format(DefaultTimeFormat),
			Version:   DefaultVersion,
		},
		BizContentRequestParams: p,
	}
	////获取签名 旧版
	//v.CommonPublicKeySign(c.AliPayPublicKey, c.AppPrivateKey, c.SignType)
	//// 转为querystring
	//str := v.toQueryString()
	//url := fmt.Sprintf("%s?%s", AlipayTradePrecreateURL, str)
	//res, _ := http.Get(url, nil, nil)
	//fmt.Println(res)

	//签名 新版
	mm := ali.M{}
	m := v.ToMap()
	for k, v := range m {
		if v != "" {
			mm[k] = v
		}
	}

	sign, _ := mm.CommonPublicKeySign(c.AliPayPublicKey, c.AppPrivateKey, "RSA2")
	mm["sign"] = sign
	qs := mm.ToQueryString(true, true)
	url := fmt.Sprintf("%s?%s", AlipayTradePrecreateURL, qs)
	res, _ := http.Get(url, nil, nil)
	fmt.Println(res)
}
