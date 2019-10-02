package alipay

import (
	"crypto/rsa"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/unliar/alipay/constans"
	"github.com/unliar/alipay/types"
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
}

// 预下单接口
func (c *Client) TradePreCreate() {
	v := types.PublicRequest{
		AppID:     c.AppID,
		Method:    constans.AlipayTradePrecreateMethodName,
		Format:    constans.DefaultFormat,
		Charset:   constans.DefaultCharset,
		SignType:  "RSA2",
		Sign:      "",
		Timestamp: time.Now().Format(constans.DefaultTimeFormat),
		Version:   constans.DefaultVersion,
		NotifyURL: "https://happysooner.com/api/v1/pay/alipay/hook",
		BizContent: types.FaceToFacePayRequest{
			OutTradeNo:  "d88da8d8ad8a8d8a8d8",
			Subject:     "测试支付",
			TotalAmount: "1.11",
		},
	}
	// 获取签名
	v.CommonPublicKeySign(c.AliPayPublicKey, c.AppPrivateKey)
	// 转为querystring
	str, err := query.Values(v)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(str)
}
