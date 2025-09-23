package mysql_model_struct

type Goods struct {
	GoodsNumber string  `gorm:"column:goods_number;primaryKey" json:"goodsNumber"` // 编号
	Name        string  `gorm:"column:name" json:"name"`
	GoodsType   string  `gorm:"column:goods_type;index" json:"goods_type"`      // 卡类型
	Amount      float64 `gorm:"column:amount;type:decimal(18,2)" json:"amount"` // 面值
	Icon        string  `gorm:"column:icon;type:varchar(255)" json:"icon"`      // 图标
	Status      int     `gorm:"column:status;type:int;default:0" json:"status"`
	CreateTime  int64   `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime  int64   `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (c Goods) TableName() string {
	return "goods"
}

// 库存表
type GoodsInventory struct {
	GoodsNumber string `gorm:"column:goods_number;primaryKey" json:"goodsNumber"`
	Inventory   int    `gorm:"column:inventory;type:int;default:0" json:"inventory"`
	CreateTime  int64  `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime  int64  `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
}

func (c GoodsInventory) TableName() string {
	return "goods_inventory"
}
