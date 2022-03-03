package model

type F_mission_file struct {
	FileId      int    `gorm:"file_id" json:"file_id"`           // 附件表ID
	MissionId   int    `gorm:"mission_id" json:"mission_id"`     // 任务表ID
	MissionUuid string `gorm:"mission_uuid" json:"mission_uuid"` // 任务表uuid
	CreateTime  string `gorm:"create_time" json:"create_time"`   // 创建时间
	FileType    string `gorm:"file_type" json:"file_type"`       // 附件类型，具体指待定
	OrderNum    int    `gorm:"order_num" json:"order_num"`       // 排序号
}
