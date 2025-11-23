package main

import "go.bug.st/serial"

type Config struct {
	Address         string // 后端监听的IP和端口，例如`127.0.0.1:8080`和`:8080`
	Password        string // 系统密码，请使用bcrypt密文
	DBFile          string // 数据库文件路径
	CleanerInterval int    // 过期数据的清理间隔，秒
	TokenDuration   int    // 登录后token的有效时间
	DefaultSize     int    // 默认的分页的单页数据量
	FilePath        string // 文件目录
	LogDuration     int    // 日志的保留时间，天
	SerialTimeout   int
	StatusDuration  int
	Serial          *serial.Mode
	SerialFile      string
	LogLevel        LogLevel   // 日志等级，例如`INFO`和`WARN`和`ERROR`
	SSL             *SSLConfig // 后端证书和私钥
	Proc            *ProcConfig
}

type SSLConfig struct {
	Enabled bool
	Cert    string
	Key     string
}

type ProcConfig struct {
	SubPowerStatus string
	SubPowerOn     string
	SubPowerOff    string
}

type LogLevel string

const (
	LogInfo  LogLevel = "INFO"
	LogWarn  LogLevel = "WARN"
	LogError LogLevel = "ERROR"
)

type Interface struct {
	SubPwrStatus string
}
