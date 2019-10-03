package alipay

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/unliar/utils/go/alipay"
	"net/url"
	"sort"
	"strings"
)

// 公共请求参数
type PublicRequestParams struct {
	AppID        string `url:"app_id" json:"app_id"`                                     //必填
	Method       string `url:"method" json:"method"`                                     // 必填
	Format       string `url:"format,omitempty" json:"format,omitempty"`                 //格式化
	Charset      string `url:"charset" json:"charset"`                                   // 字符编码
	SignType     string `url:"sign_type" json:"sign_type"`                               // 签名类型
	Sign         string `url:"sign" json:"sign"`                                         // 签名
	Timestamp    string `url:"timestamp" json:"timestamp"`                               // 时间格式
	Version      string `url:"version" json:"version"`                                   //  接口版本
	NotifyURL    string `url:"notify_url,omitempty" json:"notify_url,omitempty"`         // 回调地址
	AppAuthToken string `url:"app_auth_token,omitempty" json:"app_auth_token,omitempty"` // 应用授权
	BizContent   string `url:"biz_content" json:"biz_content"`                           // 特定请求参数
}

type BizContentRequestParams struct {
	OutTradeNo  string `json:"out_trade_no"` // 商户订单ID
	TotalAmount string `json:"total_amount"` // 总金额
	Subject     string `json:"subject"`      // 主题
}

// 请求参数
type Params struct {
	PublicRequestParams
	BizContentRequestParams
}

// 转换成 map[string]string
func (p *Params) ToMap() map[string]string {
	var m map[string]string
	f, _ := json.Marshal(p.BizContentRequestParams)
	p.PublicRequestParams.BizContent = string(f)
	str, _ := json.Marshal(p.PublicRequestParams)
	_ = json.Unmarshal(str, &m)
	return m
}
func (pub *PublicRequestParams) ToMap() map[string]string {
	var m map[string]string
	f, _ := json.Marshal(pub)
	_ = json.Unmarshal(f, &m)
	return m
}
func (pub *PublicRequestParams) toQueryString() string {
	var data []string
	for k, v := range pub.ToMap() {
		if v != "" {
			data = append(data, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(data)
	str := strings.Join(data, "&")
	return str
}

// 普通公钥签名
func (p *Params) CommonPublicKeySign(AliPayPublicKey *rsa.PublicKey, AppPrivateKey *rsa.PrivateKey, SignType string) {
	m := p.ToMap()
	var data []string
	for k, v := range m {
		if k != "sign" && v != "" {
			data = append(data, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(data)
	signStr := strings.Join(data, "&")
	fmt.Println("signStr======", signStr)
	s, cs := alipay.GetSignOpsBySignType("RSA2")
	_, err := s.Write([]byte(signStr))
	if err != nil {
		panic(err)
	}
	hashByte := s.Sum(nil)
	SignByte, err := AppPrivateKey.Sign(rand.Reader, hashByte, cs)
	if err != nil {
		panic(err)
	}
	p.Sign = url.QueryEscape(base64.StdEncoding.EncodeToString(SignByte))
}
