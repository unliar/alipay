package types

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// 公共请求参数
type PublicRequest struct {
	AppID        string                `url:"app_id" json:"app_id"`                                     //必填
	Method       string                `url:"method" json:"method"`                                     // 必填
	Format       string                `url:"format,omitempty" json:"format,omitempty"`                 //格式化
	Charset      string                `url:"charset" json:"charset"`                                   // 字符编码
	SignType     string                `url:"sign_type" json:"sign_type"`                               // 签名类型
	Sign         string                `url:"sign" json:"sign"`                                         // 签名
	Timestamp    string                `url:"timestamp" json:"timestamp"`                               // 时间格式
	Version      string                `url:"version" json:"version"`                                   //  接口版本
	NotifyURL    string                `url:"notify_url,omitempty" json:"notify_url,omitempty"`         // 回调地址
	AppAuthToken string                `url:"app_auth_token,omitempty" json:"app_auth_token,omitempty"` // 应用授权
	BizContent   *FaceToFacePayRequest `url:"biz_content" json:"biz_content"`                           // 特定请求参数
}

type FaceToFacePayRequest struct {
	OutTradeNo  string `json:"out_trade_no"` // 商户订单ID
	TotalAmount string `json:"total_amount"` // 总金额
	Subject     string `json:"subject"`      // 主题
}

// 转换成 map[string]string
func (p *PublicRequest) ToMap() map[string]string {
	var m map[string]string
	str, _ := json.Marshal(p)
	_ = json.Unmarshal(str, &m)
	return m

}

// 普通公钥签名
func (p *PublicRequest) CommonPublicKeySign(AliPayPublicKey *rsa.PublicKey, AppPrivateKey *rsa.PrivateKey) {
	m := p.ToMap()
	var data []string
	for k, v := range m {
		if k != "sign" && v != "" {
			data = append(data, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(data)
	signStr := strings.Join(data, "&")
	s := sha256.New()
	_, err := s.Write([]byte(signStr))
	if err != nil {
		panic(err)
	}
	hashByte := s.Sum(nil)
	SignByte, err := AppPrivateKey.Sign(rand.Reader, hashByte, crypto.SHA256)
	if err != nil {
		panic(err)
	}
	p.Sign = url.QueryEscape(base64.StdEncoding.EncodeToString(SignByte))
	fmt.Println(p)
}

// 证书公钥签名
func (p *PublicRequest) CertPublicKeySign() string {
	return ""
}
