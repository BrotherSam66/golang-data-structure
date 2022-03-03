package model

type F_dict_item struct {
	Id          int    `gorm:"id" json:"id"`                       // 逻辑主键，自增
	Value       int    `gorm:"value" json:"value"`                 // 字典存储值
	Label       string `gorm:"label" json:"label"`                 // 字典显示值
	DictGroupId int    `gorm:"dict_group_id" json:"dict_group_id"` // 字典分组表ID
	CreatedAt   string `gorm:"created_at" json:"created_at"`       // 创建时间
	UpdatedAt   string `gorm:"updated_at" json:"updated_at"`       // 更新时间
	DeletedAt   string `gorm:"deleted_at" json:"deleted_at"`       // 删除时间（非NULL表示软删除）
	CreateBy    int    `gorm:"create_by" json:"create_by"`         // 创建人
	UpdateBy    int    `gorm:"update_by" json:"update_by"`         // 更新人
	OrderNum    int    `gorm:"order_num" json:"order_num"`         // 排序号
}
