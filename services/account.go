package services

import (
	"time"

	"github.com/shopspring/decimal"
)

// IAccountService 账户操作
type IAccountService interface {
	// CreateAccount 创建账户
	CreateAccount(dto AccountCreatedDTO) (*AccountDTO, error)
	// Transfer 转账
	Transfer(dto AccountTransferDTO) (*AccountDTO, error)
	// StoreValue 存钱
	StoreValue(dto AccountTransferDTO) (*AccountDTO, error)
	// GetEnvelopeAccountByUserId 通过userId获取红包账户
	GetEnvelopeAccountByUserId(userId string) (*AccountDTO, error)
}

// AccountCreatedDTO 创建用户的参数
type AccountCreatedDTO struct {
	UserId       string
	Username     string
	AccountName  string
	AccountType  int8
	CurrencyCode string
	Balance      decimal.Decimal
	Status       int8
}

// AccountDTO 账户本身结构
type AccountDTO struct {
	AccountNo    string          //账户编号,账户唯一标识
	AccountName  string          //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int8            //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode string          //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       string          //用户编号, 账户所属用户
	Username     string          //用户名称
	Balance      decimal.Decimal //账户可用余额
	Status       int8            //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt    time.Time       //创建时间
	UpdatedAt    time.Time       //更新时间
}

// TradeParticipator 账户教育参与者信息
type TradeParticipator struct {
	AccountNo string
	UserId    string
	Username  string
}

// AccountTransferDTO 转账参数
type AccountTransferDTO struct {
	TradeNo string
	// 交易主体
	TradeBody TradeParticipator
	// 交易目标
	TradeTarget TradeParticipator
	Amount      decimal.Decimal
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Decs        string
}

// 账户流水
type AccountLogDTO struct {
	LogNo           string          // 流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string          // 交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string          // 账户编号 账户ID
	TargetAccountNo string          // 账户编号 账户ID
	UserId          string          // 用户编号
	Username        string          // 用户名称
	TargetUserId    string          // 目标用户编号
	TargetUsername  string          // 目标用户名称
	Amount          decimal.Decimal // 交易金额,该交易涉及的金额
	Balance         decimal.Decimal // 交易后余额,该交易后的余额
	ChangeType      ChangeType      // 流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      ChangeFlag      // 交易变化标识：-1 出账 1为进账，枚举
	Status          int             // 交易状态：
	Decs            string          // 交易描述
	CreatedAt       time.Time       // 创建时间
}