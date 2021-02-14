package model

import "time"

// Site describe a model which stores the sites need to be monitored
type Site struct {
	ID         uint      `json:"id"   gorm:"primaryKey"`
	Tag        string    `json:"tag"  gorm:"type:char(36);unique"`
	Name       string    `json:"name" gorm:"type:vchar(63)"`
	Public     int       `json:"public" gorm:"default:0"`
	LastCheck  time.Time `json:"last_check"`
	FirstCheck time.Time `json:"first_check"`
}

// Group describe a model which stores the group info
type Group struct {
	ID   uint   `json:"id"   gorm:"primaryKey"`
	Link string `json:"link" gorm:"type:text"`
	Site []Site `gorm:"many2many:site_group"`
}

// site_group table

// Condition describe a model which list the warnning threshold of Site
type Condition struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Site         []Site  `gorm:"many2many:site_condition"`
	IsEnable     uint    `json:"is_enable"`
	Link         string  `json:"link"`
	ResType      uint    `json:"res_type"`
	ResThreshold float32 `json:"res_threshold"`
}

// site_condition table

// Notification describe how to send message to admin
type Notification struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Info string `json:"info" gorm:"type:vchar(1023)"`
	Site []Site `gorm:"many2many:site_notification"`
}

// site_notification table

// Log describe what happened
type Log struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}
