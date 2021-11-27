package model

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
)

type GqaPluginExampleNews struct {
	UpdatedByUser *system.SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *system.SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	Title   string `json:"title" gorm:"comment:新闻标题;not null;index"`
	Content string `json:"content" gorm:"comment:新闻内容;not null;"`
}
