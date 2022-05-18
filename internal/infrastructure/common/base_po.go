package common

import "time"

type BasePO struct {
	Status     uint8     `json:"status"`
	CreateBy   uint      `json:"-"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy   uint      `json:"-"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	Remarks    string    `json:"remarks"`
}
