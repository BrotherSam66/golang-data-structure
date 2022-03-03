package model

type F_mission_file_his struct {
	FileHisId   int    `gorm:"file_his_id" json:"file_his_id"`     // 逻辑主键，自增
	FileId      int    `gorm:"file_id" json:"file_id"`             // 任务附件关联表对应字段快照
	MissionId   int    `gorm:"mission_id" json:"mission_id"`       // 任务附件关联表对应字段快照
	MissionUuid string `gorm:"mission_uuid" json:"mission_uuid"`   // 任务附件关联表对应字段快照
	CreateTime  string `gorm:"create_time" json:"create_time"`     // 任务附件关联表对应字段快照
	FileType    string `gorm:"file_type" json:"file_type"`         // 任务附件关联表对应字段快照
	OrderNum    int    `gorm:"order_num" json:"order_num"`         // 任务附件关联表对应字段快照
	HistoryTime string `gorm:"history_time" json:"history_time"`   // 创建时间
	ActionLogId int    `gorm:"action_log_id" json:"action_log_id"` // 任务变更日志表ID
}
