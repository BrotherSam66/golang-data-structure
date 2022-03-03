package model

type F_mission_his struct {
	HisId                int    `gorm:"his_id" json:"his_id"`                                 // 任务快照表
	Id                   int    `gorm:"id" json:"id"`                                         // 任务表对应字段快照
	Uuid                 string `gorm:"uuid" json:"uuid"`                                     // 任务表对应字段快照
	Path                 string `gorm:"path" json:"path"`                                     // 任务表对应字段快照
	Name                 string `gorm:"name" json:"name"`                                     // 任务表对应字段快照
	Content              string `gorm:"content" json:"content"`                               // 任务表对应字段快照
	Type                 int    `gorm:"type" json:"type"`                                     // 任务表对应字段快照
	DeptId               int    `gorm:"dept_id" json:"dept_id"`                               // 任务表对应字段快照
	UserId               int    `gorm:"user_id" json:"user_id"`                               // 任务表对应字段快照
	StartTime            string `gorm:"start_time" json:"start_time"`                         // 任务表对应字段快照
	RequiredComplateTime string `gorm:"required_complate_time" json:"required_complate_time"` // 任务表对应字段快照
	CreateTime           string `gorm:"create_time" json:"create_time"`                       // 任务表对应字段快照
	CreateBy             int    `gorm:"create_by" json:"create_by"`                           // 任务表对应字段快照
	UpdateTime           string `gorm:"update_time" json:"update_time"`                       // 任务表对应字段快照
	UpdateBy             int    `gorm:"update_by" json:"update_by"`                           // 任务表对应字段快照
	Level                int    `gorm:"level" json:"level"`                                   // 任务表对应字段快照
	State                int    `gorm:"state" json:"state"`                                   // 任务表对应字段快照
	Priority             int    `gorm:"priority" json:"priority"`                             // 任务表对应字段快照
	ParentId             string `gorm:"parent_id" json:"parent_id"`                           // 任务表对应字段快照
	OrderNum             int    `gorm:"order_num" json:"order_num"`                           // 任务表对应字段快照
	Remark               string `gorm:"remark" json:"remark"`                                 // 任务表对应字段快照
	HistoryTime          string `gorm:"history_time" json:"history_time"`                     // 创建时间
	ActionLogId          int    `gorm:"action_log_id" json:"action_log_id"`                   // 任务变更日志表ID
}
