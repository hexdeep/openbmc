package main

import "time"

type Image struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Filename  string    `gorm:"type:VARCHAR(255);not null;comment:文件名"`
	Size      uint64    `gorm:"not null;comment:大小"`
}
