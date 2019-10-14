package alipay

import (
	"encoding/json"
	"fmt"
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

// 商家请求参数
type BizContentRequestParams struct {
	OutTradeNo  string `json:"out_trade_no,omitempty"` // 商户订单ID
	TotalAmount string `json:"total_amount,omitempty"` // 总金额
	Subject     string `json:"subject,omitempty"`      // 主题
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

//------------****响应参数********** ----------------------------------------
type PublicResponse struct {
	Code    string `json:"code"`     // 网关返回代码
	Msg     string `json:"msg"`      // 网关错误信息
	SubCode string `json:"sub_code"` // 业务返回代码
	SumMsg  string `json:"sum_msg"`  // 业务返回描述
}

type TradePreCreateResponse struct {
	AlipayTradePrecreateResponse struct {
		PublicResponse
		OutTradeNo string `json:"out_trade_no,omitempty"` // 商家订单号
		QrCode     string `json:"qr_code,omitempty"`      // 商家订单号
	} `json:"alipay_trade_precreate_response"`
}

// 交易结果查询
type TradeQueryResponse struct {
	AlipayTradeQueryResponse struct {
		PublicResponse
		TradeStatus string `json:"trade_status"`  // 交易状态
		TotalAmount string `json:"total_amount"`  // 总金额
		BuyerUserID string `json:"buyer_user_id"` // 买家金额
	} `json:"alipay_trade_query_response"`
}

// 原始请求字符转换成签名字符串
func SignRawStrConvert(rawStr string) string {
	var signStrArr []string
	// 签名验证
	str, _ := url.PathUnescape(rawStr)
	v, _ := url.ParseQuery(str)
	for k, item := range v {
		if k != "sign" && k != "sign_type" && len(item) > 0 {
			fmt.Println(k, item[0])
			signStrArr = append(signStrArr, fmt.Sprintf("%s=%s", k, item[0]))
		}
	}
	sort.Strings(signStrArr)
	return strings.Join(signStrArr, "&")
}
