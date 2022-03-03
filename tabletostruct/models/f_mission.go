package model

type F_mission struct {
	Id                   int    `gorm:"id" json:"id"`                                         // 逻辑主键，自增
	Uuid                 string `gorm:"uuid" json:"uuid"`                                     // 任务唯一编号,uuid或者时间日期编号
	Path                 string `gorm:"path" json:"path"`                                     // 任务层级路径
	Name                 string `gorm:"name" json:"name"`                                     // 任务名称
	Content              string `gorm:"content" json:"content"`                               // 任务内容
	Type                 int    `gorm:"type" json:"type"`                                     // 任务类型 取字典表
	DeptId               int    `gorm:"dept_id" json:"dept_id"`                               // 责任部门
	UserId               int    `gorm:"user_id" json:"user_id"`                               // 责任人
	StartTime            string `gorm:"start_time" json:"start_time"`                         // 任务开始时间
	RequiredComplateTime string `gorm:"required_complate_time" json:"required_complate_time"` // 要求完成时间
	CreateTime           string `gorm:"create_time" json:"create_time"`                       // 创建时间
	CreateBy             int    `gorm:"create_by" json:"create_by"`                           // 创建人
	UpdateTime           string `gorm:"update_time" json:"update_time"`                       // 最后更新时间
	UpdateBy             int    `gorm:"update_by" json:"update_by"`                           // 最后更新人
	Level                int    `gorm:"level" json:"level"`                                   // 任务层级
	State                int    `gorm:"state" json:"state"`                                   // 任务状态，具体值待定
	Priority             int    `gorm:"priority" json:"priority"`                             // 任务优先级 取字典表
	ParentId             string `gorm:"parent_id" json:"parent_id"`                           // 父任务uuid字段
	OrderNum             int    `gorm:"order_num" json:"order_num"`                           // 排序号
	Remark               string `gorm:"remark" json:"remark"`                                 // 备注
}
