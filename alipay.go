package alipay

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	ali "github.com/unliar/utils/go/alipay"
	"github.com/unliar/utils/go/http"
	"time"
)

type Client struct {
	AppID           string         // 应用id
	NotifyURL       string         // 回调接口
	SignType        string         // 签名类型
	EndpointURL     string         // 创建交易地址 必须要有
	AliPayPublicKey *rsa.PublicKey // 支付宝公钥
	// 普通公钥签名可用
	AppPrivateKey *rsa.PrivateKey // 应用私钥
	// 以下公钥证书可以
	AppPublicKeyCert    string // app公钥证书
	AliPayPublicKeyCert string // 支付宝公钥证书
	AliPayRootCert      string // 支付宝根证书
	ReturnURL           string // wap 支付成功返回地址
}

// NewClient 新实例化
func NewClient(c Client) *Client {
	return &Client{
		AppID:               c.AppID,
		NotifyURL:           c.AppID,
		SignType:            c.SignType,
		EndpointURL:         c.EndpointURL,
		AliPayPublicKey:     c.AliPayPublicKey,
		AppPrivateKey:       c.AppPrivateKey,
		AppPublicKeyCert:    c.AppPublicKeyCert,
		AliPayPublicKeyCert: c.AliPayPublicKeyCert,
		AliPayRootCert:      c.AliPayRootCert,
		ReturnURL:           c.ReturnURL,
	}
}

// wap 支付设置成功支付回跳地址
func (c *Client) SetReturnURL(str string) {
	c.ReturnURL = str
}

// 预下单接口
func (c *Client) TradePreCreate(p BizContentRequestParams) (*TradePreCreateResponse, error) {
	var tpr TradePreCreateResponse
	res, err := c.DoRequest(AlipayTradePrecreateMethodName, p, false)
	if err != nil {
		return &tpr, err
	}
	_ = json.Unmarshal(res, &tpr)
	return &tpr, nil
}

// 网站支付下单接口
func (c *Client) TradePagePay(p BizContentRequestParams) (string, error) {
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	res, err := c.DoRequest(AlipayTradePagePay, p, true)
	return string(res), err
}

// wap下单接口
func (c *Client) TradeWapPay(p BizContentRequestParams, ReturnUrl string) (string, error) {
	p.ProductCode = "QUICK_WAP_WAY"
	c.SetReturnURL(ReturnUrl)
	res, err := c.DoRequest(AlipayTradeWapPay, p, true)
	return string(res), err
}

// 查询订单接口
func (c *Client) TradeQuery(p BizContentRequestParams) (*TradeQueryResponse, error) {
	var tpr TradeQueryResponse
	res, err := c.DoRequest(AlipayTradeQueryMethodName, p, false)
	if err != nil {
		return &tpr, err
	}
	_ = json.Unmarshal(res, &tpr)
	return &tpr, nil
}

// 撤销订单接口
func (c *Client) TradeCancel(p BizContentRequestParams) (*TradeCancelResponse, error) {
	var tpr TradeCancelResponse
	res, err := c.DoRequest(AlipayTradeCancelMethodName, p, false)
	if err != nil {
		return &tpr, err
	}
	_ = json.Unmarshal(res, &tpr)
	return &tpr, nil
}

// 交易退款
func (c *Client) TradeRefund(p BizContentRequestParams) (*TradeCancelResponse, error) {
	var tpr TradeCancelResponse
	res, err := c.DoRequest(AlipayTradeRefundMethodName, p, false)
	if err != nil {
		return &tpr, err
	}
	_ = json.Unmarshal(res, &tpr)
	return &tpr, nil
}

// 通用请求接口
func (c *Client) DoRequest(method string, p BizContentRequestParams, fromBrowser bool) ([]byte, error) {
	v := Params{
		PublicRequestParams: PublicRequestParams{
			AppID:     c.AppID,
			Method:    method,
			Format:    DefaultFormat,
			Charset:   DefaultCharset,
			SignType:  c.SignType,
			NotifyURL: c.NotifyURL,
			Timestamp: time.Now().Format(DefaultTimeFormat),
			Version:   DefaultVersion,
			ReturnURL: c.ReturnURL,
		},
		BizContentRequestParams: p,
	}

	//签名 新版
	mm := ali.M{}
	m := v.ToMap()
	for k, v := range m {
		if v != "" {
			mm[k] = v
		}
	}
	sign, _ := mm.CommonPublicKeySign(c.AliPayPublicKey, c.AppPrivateKey, c.SignType)
	mm["sign"] = sign
	qs := mm.ToQueryString(true, true)
	url := fmt.Sprintf("%s?%s", c.EndpointURL, qs)
	if fromBrowser {
		fmt.Println("当前请求不走服务器请求,只返回一个网址,请把网址填入你的网页iframe src")
		return []byte(url), nil
	}
	res, err := http.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 检查签名
func (c *Client) CheckSign(signReqStr string) bool {
	str, sign := SignRawStrConvert(signReqStr)
	signByte, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println(" base64 err==>", err)
		return false
	}
	hash, cto := ali.GetSignOpsBySignType(c.SignType)
	_, err = hash.Write([]byte(str))
	if err != nil {
		return false
	}
	hashByte := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(c.AliPayPublicKey, cto, hashByte, signByte)
	if err != nil {
		return false
	}
	return true
}
