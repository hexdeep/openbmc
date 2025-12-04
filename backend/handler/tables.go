package handler

import "time"

var Tables = []any{
	new(Token),
	new(Log),
}

type Token struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	ExpiresAt time.Time `gorm:"not null;comment:过期时间"`
}

type Loga struct {
	ID        uint      `gorm:"primarykey;comment:标识"`
	CreatedAt time.Time `gorm:"not null;comment:创建时间"`
	ClientIP  string    `gorm:"type:VARCHAR(50);not null;comment:客户端地址"`
	Method    string    `gorm:"type:VARCHAR(10);not null;comment:请求方法"`
	Path      string    `gorm:"type:VARCHAR(30);not null;comment:路径"`
	Status    int       `gorm:"not null;comment:状态码"`
}
