package bean

import (
	"github.com/jinzhu/gorm"
)

// Version 版本管理表
type Version struct {
	gorm.Model
	VersionBase
}

// VersionBase VersionBase
type VersionBase struct {
	Version    string `gorm:"not null" form:"new_version" json:"new_version"` // app 版本号
	Type       string `gorm:"not null" form:"type" json:"type"`               // 1 安卓, 2 ios
	ApkURL     string `gorm:"not null" form:"apk_url" json:"apk_url"`         // 安装包下载地址
	UpdateLog  string `gorm:"not null" form:"update_log" json:"update_log"`   // 更新内容
	Delta      string `form:"delta" json:"delta"`                             // Delta
	NewMd5     string `form:"new_md5" json:"new_md5"`                         // md5
	TargetSize string `form:"target_size" json:"target_size"`                 // TargetSize
}
