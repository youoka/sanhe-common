package mysql_model_struct

type RechargeCategory struct {
	ID         uint           `gorm:"primaryKey" json:"id"`                          // id
	Name       string         `gorm:"type:varchar(100)" json:"name"`                 // 商品名称
	Img        string         `gorm:"type:varchar(255)" json:"img"`                  // 图片地址
	Mode       int            `gorm:"default:0" json:"mode"`                         // mode
	Sort       int            `gorm:"default:0" json:"sort"`                         // 新增字段
	Color      string         `gorm:"type:varchar(50)" json:"color"`                 // 新增字段
	Children   []RechargeCard `gorm:"foreignKey:Type;references:ID" json:"children"` // 关联卡列表
	CreateTime int64          `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime int64          `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (c RechargeCategory) TableName() string {
	return "recharge_category"
}

type RechargeCard struct {
	ID         uint               `gorm:"primaryKey" json:"id"`                                      // id
	Name       string             `gorm:"type:varchar(100)" json:"name"`                             // 商品名称
	Img        string             `gorm:"type:varchar(255)" json:"img"`                              // 图片地址
	Discount   string             `gorm:"type:varchar(10)" json:"discount"`                          // 折扣
	Lizi       string             `gorm:"type:varchar(255)" json:"lizi"`                             // 卡号信息
	Mark       int                `gorm:"default:0" json:"mark"`                                     // mark
	Mode       int                `gorm:"default:0" json:"mode"`                                     // mode
	Nocode     int                `gorm:"default:0" json:"nocode"`                                   // nocode
	Note       string             `gorm:"type:text" json:"note"`                                     // 商品规则说明
	Repair     int                `gorm:"default:0" json:"repair"`                                   // repair
	Route      string             `gorm:"type:varchar(50)" json:"route"`                             // route
	Shide      int                `gorm:"default:0" json:"shide"`                                    // shide
	Type       int                `gorm:"default:1" json:"type"`                                     // type
	Warning    int                `gorm:"default:0" json:"warning"`                                  // warning
	Auto       int                `gorm:"default:1" json:"auto"`                                     // auto
	Canal      int                `gorm:"default:0" json:"canal"`                                    // canal
	Start      int                `gorm:"default:1" json:"start"`                                    // start
	ListItems  []RechargeCardItem `gorm:"foreignKey:CardID;constraint:OnDelete:CASCADE" json:"list"` // 面额列表
	CreateTime int64              `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime int64              `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (c RechargeCard) TableName() string {
	return "recharge_card"
}

type RechargeCardItem struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	CardID     uint   `gorm:"index" json:"card_id"`          // 外键，关联 RechargeCard
	LType      int    `gorm:"default:0" json:"ltype"`        // ltype
	Name       string `gorm:"type:varchar(50)" json:"name"`  // 面额名称
	Note       string `gorm:"type:varchar(255)" json:"note"` // 备注
	Price      string `gorm:"type:varchar(20)" json:"price"` // 价格
	CreateTime int64  `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime int64  `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (c RechargeCardItem) TableName() string {
	return "recharge_card_item"
}
