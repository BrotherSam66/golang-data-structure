package model

type F_mission_action_log struct {
	Id         int    `gorm:"id" json:"id"`                   // 逻辑主键 自增
	UserId     int    `gorm:"user_id" json:"user_id"`         // 操作用户ID
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	ActionType string `gorm:"action_type" json:"action_type"` // 操作类型，具体值待定，新增，修改，删除，状态变更等
	Content    string `gorm:"content" json:"content"`         // 一些需要保存的额外信息
}
