// 这里是支付宝常量接口定义
package alipay

// 统一收单线下请求地址
const AlipayTradeEndpointURL = "https://openapi.alipay.com/gateway.do"

// 统一收单线下交易预创 接口名称
const AlipayTradePrecreateMethodName = "alipay.trade.precreate"

// 统一收单线下交易查询  接口名称
const AlipayTradeQueryMethodName = "alipay.trade.query"

// 统一收单线下撤销  接口名称
const AlipayTradeCancelMethodName = "alipay.trade.cancel"

// 统一收单同步退款  接口名称
const AlipayTradeRefundMethodName = "alipay.trade.refund"

// https://opendocs.alipay.com/pre-open/270/105900  2020-03-07 新增电脑网站支付接口
// 统一收单下单并支付页面接口
const AlipayTradePagePay = "alipay.trade.page.pay"

// 统一收到wap支付接口
const AlipayTradeWapPay = "alipay.trade.wap.pay"

// 格式
const DefaultFormat = "JSON"

// 编码
const DefaultCharset = "utf-8"

// 版本
const DefaultVersion = "1.0"

// 时间戳格式
const DefaultTimeFormat = "2006-01-02 15:04:05"
