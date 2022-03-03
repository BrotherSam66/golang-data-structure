package model

type F_dict_group struct {
	Id         int    `gorm:"id" json:"id"`                   // 逻辑主键，自增
	Name       string `gorm:"name" json:"name"`               // 字典组名称，如任务类型，任务优先级
	Code       string `gorm:"code" json:"code"`               // 字典组唯一编码，用于代码中的查询
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	UpdateTime string `gorm:"update_time" json:"update_time"` // 更新时间
	CreateBy   int    `gorm:"create_by" json:"create_by"`     // 创建人
	UpdateBy   int    `gorm:"update_by" json:"update_by"`     // 更新人
}
