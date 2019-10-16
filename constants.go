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

// 格式
const DefaultFormat = "JSON"

// 编码
const DefaultCharset = "utf-8"

// 版本
const DefaultVersion = "1.0"

// 时间戳格式
const DefaultTimeFormat = "2006-01-02 15:04:05"
