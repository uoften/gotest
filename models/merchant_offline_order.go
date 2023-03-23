package sun9ShouMi_offline_order_db

import "gopkg.in/mgo.v2/bson"

var (
	OffLineOrderDbName         = "sun9db_shoumi_order_"
	OffLineOrderCollectionName = "order_"
)

type OrderInfo struct {
	Id                           bson.ObjectId       `bson:"_id"`                                    // 订单ID
	ExtOrderId                   string              `bson:"extOrderId"`                             // 外部订单号
	ChannelOrderId               string              `bson:"channelOrderId"`                         // 通道订单号
	ThirdOrderId                 string              `bson:"thirdOrderId"`                           // 银行，微信，支付宝侧订单号
	TradeTime                    bson.MongoTimestamp `bson:"tradeTime"`                              // 交易发起时间
	PayedTime                    int64               `bson:"payedTime"`                              // 支付时间完成时间
	TransAmount                  int                 `bson:"transAmount"`                            // 交易金额 单位 分
	Fee                          int                 `bson:"fee"`                                    // 交易费率 单位：万分之一
	Charge                       int                 `bson:"charge"`                                 // 交易手续费
	SettleAmount                 int                 `bson:"settleAmount"`                           // 结算金额
	SettleStatus                 int                 `bson:"settleStatus"`                           //结算状态：1-待支付 2-支付成功 3-输密支付中 4-支付失败 5-订单关闭 6-退款中 7-退款成功 8-退款失败 9-冻结
	PayWay                       int                 `bson:"payWay"`                                 //支付方式 1-微信 2-支付宝 3-POS刷卡 4-银联扫码 5-云闪付 6-京东钱包
	PaySource                    int                 `bson:"paySource"`                              //支付类型，来源（立牌，扫码枪，插件）
	TransType                    int                 `bson:"transType"`                              //预授权 刷卡交易状态  1-消费 2-消费冲正 3-消费撤销 4-冻结
	OrderType                    int                 `bson:"orderType"`                              // 1-收款订单  2-退款订单 3-会员储值订单
	MchtDiscountType             int                 `bson:"mchtDiscountType"`                       //商户优惠类型 1-标准类 2-优惠类
	PayBank                      string              `bson:"payBank"`                                //付款银行
	CardType                     int                 `bson:"cardType"`                               //付款卡类型 1-借记卡  2-信用卡
	PosNo                        string              `bson:"posNo"`                                  //终端号 刷卡支付字段
	RefNo                        string              `bson:"refNo"`                                  //参考号 刷卡支付字段
	VoucherNo                    string              `bson:"voucherNo"`                              //凭证号  刷卡支付字段
	PromotionTag                 int                 `bson:"promotionTag"`                           //活动类型 0-无活动  1-蓝海优惠 2-绿洲优惠
	DiscountType                 int                 `bson:"discountType"`                           //代金券类型 1-免充金， 2-预冲金
	DiscountAmount               int                 `bson:"discountAmount"`                         //订单优惠金额
	UserOpenId                   string              `bson:"userOpenId"`                             // 用户OpenId 或 支付宝ID
	RefundOriginalOrderId        string              `bson:"refundOriginalOrderId"`                  //退款的原订单号
	RefundChannelOriginalOrderId string              `bson:"refundChannelOriginalOrderId"`           // 通道退款原订单号
	ChannelReceptId              string              `bson:"channelReceptId"`                        // 通道收单号
	ChannelType                  int                 `bson:"channelType"`                            // 通道类型 1-乐刷 2-富有 3-威富通
	MchtReceiptId                string              `bson:"mchtReceiptId"`                          // 商户收单号
	MchtMainId                   string              `bson:"mchtMainId"`                             // 商户主体ID
	AgentId                      string              `bson:"agentId"`                                // 代理商ID
	DevPartnerId                 string              `bson:"devPartnerId"`                           // 开发商ID
	DevSn                        string              `bson:"devSn"`                                  // 交易设备SN
	StoreID                      string              `bson:"storeID"`                                // 门店ID
	StoreName                    string              `bson:"storeName"`                              // 门店名称
	OprId                        string              `bson:"oprId"`                                  // 操作员ID
	OprName                      string              `bson:"oprName"`                                // 操作员名称
	ErrMsg                       string              `bson:"errMsg"`                                 // 订单失败原因
	Coupon                       int                 `bson:"coupon"`                                 // 支付宝红包优惠金
	CheckStatus                  int32               `bson:"checkStatus"`                            //对账状态 0-未对账  1-已对账  2-对账补充订单 3-对账延期订单 4-对账失败订单
	CheckTime                    bson.MongoTimestamp `bson:"checkTime"`                              //对账时间
	ChannelCharge                int64               `bson:"channel_charge"`                         //通道手续费
	ChannelFee                   int64               `bson:"channel_fee"`                            //通道费率
	ChannelSettleAmount          int64               `bson:"channel_settle_amount"`                  //通道结算金额
	ChannelTransAmount           int64               `bson:"channel_trans_amount"`                   //通道交易金额
	MerchantName                 string              `bson:"merchantName"`                           //商户名称
	StoreCityCode                string              `bson:"store_city_code",json:"store_city_code"` //门店城市code
	StoreCityName                string              `bson:"store_city_name",json:"store_city_name"` //门店城市名称
	StoreProvenceCode            string              `bson:"store_provence_code",json:"store_provence_code"`
	StoreProvenceName            string              `bson:"store_provence_name",json:"store_provence_name"`
	StoreAreaCode                string              `bson:"store_area_code",json:"store_area_code"`
	StoreAreaName                string              `bson:"store_area_name",json:"store_area_name"`
	UserId                       string              `bson:"user_id"`                                 //用户id
}
