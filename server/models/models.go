package models

import (
	"gorm.io/gorm"
	"time"
)

// Movie 影视资源模型
type Movie struct {
	gorm.Model
	Title       string    `json:"title" gorm:"not null"` // 影视标题
	Description string    `json:"description"`            // 影视描述
	Cover       string    `json:"cover"`                 // 封面图片URL
	Year        int       `json:"year"`                  // 发行年份
	Director    string    `json:"director"`              // 导演
	Actors      string    `json:"actors"`                // 主演
	Type        string    `json:"type"`                  // 类型（电影/电视剧）
	Region      string    `json:"region"`                // 地区
	Links       []Link    `json:"links"`                 // 网盘链接
	Tags        []Tag     `json:"tags" gorm:"many2many:movie_tags"` // 标签
	Views       int       `json:"views" gorm:"default:0"` // 浏览次数
	CreatedBy   uint      `json:"created_by"`            // 创建者ID
}

// Link 网盘链接模型
type Link struct {
	gorm.Model
	MovieID     uint      `json:"movie_id"`              // 关联的影视资源ID
	Platform    string    `json:"platform" gorm:"not null"` // 网盘平台（百度网盘/阿里云盘等）
	URL         string    `json:"url" gorm:"not null"`    // 分享链接
	Password    string    `json:"password"`              // 提取密码
	ExpireTime  time.Time `json:"expire_time"`           // 过期时间
	Status      int       `json:"status" gorm:"default:1"` // 链接状态（1:有效 0:无效）
}

// Tag 标签模型
type Tag struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique;not null"` // 标签名称
	Description string    `json:"description"`                  // 标签描述
	Movies      []Movie   `json:"movies" gorm:"many2many:movie_tags"` // 关联的影视资源
}

// User 用户模型
type User struct {
	gorm.Model
	Username    string    `json:"username" gorm:"unique;not null"` // 用户名
	Password    string    `json:"password" gorm:"not null"`        // 密码（加密存储）
	Nickname    string    `json:"nickname"`                        // 昵称
	Email       string    `json:"email" gorm:"unique"`            // 邮箱
	Avatar      string    `json:"avatar"`                         // 头像URL
	Role        int       `json:"role" gorm:"default:1"`         // 角色（1:普通用户 2:管理员）
	Status      int       `json:"status" gorm:"default:1"`       // 状态（1:正常 0:禁用）
	LastLogin   time.Time `json:"last_login"`                     // 最后登录时间
}