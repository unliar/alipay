package requestParams

// 公共请求参数
type PublicRequest struct {
	AppID        string               `json:"app_id"`                   //必填
	Method       string               `json:"method"`                   // 必填
	Format       string               `json:"format,omitempty"`         //格式化
	Charset      string               `json:"charset"`                  // 字符编码
	SignType     string               `json:"sign_type"`                // 签名类型
	Sign         string               `json:"sign"`                     // 签名
	Timestamp    string               `json:"timestamp"`                // 时间格式
	Version      string               `json:"version"`                  //  接口版本
	NotifyURL    string               `json:"notify_url,omitempty"`     // 回调地址
	AppAuthToken string               `json:"app_auth_token,omitempty"` // 应用授权
	BizContent   FaceToFacePayRequest `json:"biz_content"`              // 特定请求参数
}

type FaceToFacePayRequest struct {
	OutTradeNo  string `json:"out_trade_no"` // 商户订单ID
	TotalAmount string `json:"total_amount"` // 总金额
	Subject     string `json:"subject"`      // 主题
}

// 普通公钥签名
func (p *PublicRequest) CommonPublicKeySign() {

}

// 证书公钥签名
func (p *PublicRequest) CertPublicKeySign() string {
	return ""
}
