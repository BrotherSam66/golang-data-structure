package model

type F_mission_user struct {
	UserId      int    `gorm:"user_id" json:"user_id"`           // 用户id
	MissionId   int    `gorm:"mission_id" json:"mission_id"`     // 任务表ID
	Type        string `gorm:"type" json:"type"`                 // 关联类型，具体值待定，目前设置只有配合用户
	MissionUuid string `gorm:"mission_uuid" json:"mission_uuid"` // 任务表uuid
	CreateTime  string `gorm:"create_time" json:"create_time"`   // 创建时间
}
