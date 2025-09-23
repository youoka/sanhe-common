package mysql_model_struct

import "time"

// BalanceLog 余额变动记录表
type BalanceLog struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Account     string    `gorm:"column:account;type:varchar(64);index" json:"account"`
	Amount      float64   `gorm:"column:amount;type:decimal(18,2)" json:"amount"`
	Balance     float64   `gorm:"column:balance;type:decimal(18,2)" json:"balance"`               // 变动后余额
	LogType     int       `gorm:"column:log_type;type:int" json:"log_type"`                       // 1:充值 2:消费 3:退款等
	OrderNumber string    `gorm:"column:order_number;type:varchar(255);index" json:"orderNumber"` // 关联订单号
	Remark      string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime;index" json:"createTime"`
}

func (b BalanceLog) TableName() string {
	return "balance_logs"
}
