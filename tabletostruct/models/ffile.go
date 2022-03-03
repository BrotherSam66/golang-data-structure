package model

type F_file struct {
	Id         int    `gorm:"id" json:"id"`                   // 逻辑主键 自增
	FileName   string `gorm:"file_name" json:"file_name"`     // 附件名称
	Path       string `gorm:"path" json:"path"`               // 文件存放路径
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	CreateBy   int    `gorm:"create_by" json:"create_by"`     // 上传用户
}
