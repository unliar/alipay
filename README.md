# alipay

[![CircleCI](https://circleci.com/gh/unliar/alipay/tree/master.svg?style=svg)](https://circleci.com/gh/unliar/alipay/tree/master)

支付宝付款sdk,目前只实现了普通公钥模式的部分接口

> https://docs.open.alipay.com/194/105072
## API
1. TradePreCreate 预下单接口 - 已线上测试通过
2. TradeQuery 查询订单接口 - 已线上测试通过
3. CheckSign 签名验证接口 - 已线上测试通过
4. TradeCancel 交易撤销接口 - 线上未测试
5. TradeRefund  交易退款 - 未完成
6. TradePagePay 电脑网站支付下单接口 - 已线上测试
7. TradeWapPay  手机网站支付下单接口 - 已线上测试
> 6 7 返回一个奇怪的网址 只要把它嵌入网站的 iframe 标签就会自动跳转到支付宝收银台
## Demo

> 使用案例见 example 文件夹

## Test 
> to do 

