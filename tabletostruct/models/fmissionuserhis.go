package model

type F_mission_user_his struct {
	Id          int    `gorm:"id" json:"id"`                       // 逻辑主键，自增
	UserId      int    `gorm:"user_id" json:"user_id"`             // 任务用户关联表对应字段快照
	MissionId   int    `gorm:"mission_id" json:"mission_id"`       // 任务用户关联表对应字段快照
	Type        string `gorm:"type" json:"type"`                   // 任务用户关联表对应字段快照
	MissionUuid string `gorm:"mission_uuid" json:"mission_uuid"`   // 任务用户关联表对应字段快照
	CreateTime  string `gorm:"create_time" json:"create_time"`     // 任务用户关联表对应字段快照
	HistoryTime string `gorm:"history_time" json:"history_time"`   // 创建时间
	ActionLogId int    `gorm:"action_log_id" json:"action_log_id"` // 任务变更日志表ID
}
