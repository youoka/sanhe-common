package mysql_model_struct

import "time"

type User struct {
	Account        string    `gorm:"column:account;primary_key;type:varchar(64)" json:"account"`
	Password       string    `gorm:"column:password;type:varchar(255)" json:"password"`
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(32);uniqueIndex" json:"phoneNumber"`
	ParentAccount  string    `gorm:"column:parent_account;type:varchar(64);index" json:"parentAccount"` // 直接上级账号
	CreateTime     time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	Balance        float64   `gorm:"column:balance;type:decimal(18,2);default:0" json:"balance"`
	CommissionRate float64   `gorm:"column:commission_rate;type:decimal(18,2);default:0" json:"commissionRate"` // 佣金比例
	Status         int       `gorm:"column:status;type:int;default:0" json:"status"`                            // 0 normal, 1 disabled
}

func (u User) TableName() string {
	return "users"
}

type AuthCode struct {
	Account     string `gorm:"column:account;primary_key;type:varchar(64)" json:"account"`
	QrSecretKey string `gorm:"column:qr_secret_key;type:varchar(255)" json:"qrSecretKey"`
}

func (q AuthCode) TableName() string {
	return "auth_codes"
}

// 用户代理关系表
type UserHierarchy struct {
	Account       string    `gorm:"column:account;primary_key;type:varchar(64)" json:"account"`
	ParentAccount string    `gorm:"column:parent_account;primary_key;type:varchar(64)" json:"parentAccount"` // 上级账号
	Level         int       `gorm:"column:level;type:int" json:"level"`                                      // 层级(1-直接上级, 2-二级上级...)
	Rates         float64   `gorm:"column:rates;type:decimal(18,2);default:0" json:"rates"`
	CreateTime    time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
}

func (u UserHierarchy) TableName() string {
	return "user_hierarchies"
}
