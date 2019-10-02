package alipay

import (
	"crypto/rsa"
	"fmt"
	"github.com/unliar/alipay/constans"
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
	url := constans.AlipayTradePrecreateURL
	fmt.Println(url)
}
