package alipay

import (
	"crypto/rsa"
	"fmt"
	"github.com/unliar/alipay/constans"
)

type Client struct {
	AppID          string          // 应用id
	AppPrivateKey  *rsa.PrivateKey // 应用私钥
	AliayPublicKey *rsa.PublicKey  // 支付宝公钥
}

// 预下单接口
func (c *Client) TradePreCreate() {
	url := constans.AlipayTradePrecreateURL
	fmt.Println(url)
}
