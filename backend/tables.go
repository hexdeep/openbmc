package main

import "time"

var tables = []any{
	new(Token),
	new(Log),
}

type Token struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	ExpiresAt time.Time `gorm:"not null;comment:过期时间"`
}

type Log struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	Method    string    `gorm:"type:VARCHAR(10);not null;comment:请求方法"`
	Path      string    `gorm:"type:VARCHAR(30);not null;comment:路径"`
	Status    int       `gorm:"not null;comment:状态码"`
}
