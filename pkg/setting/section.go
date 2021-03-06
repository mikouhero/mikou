package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogMaxSize      int
	LogMaxAge       int
	LogFileExt      string
	AuthPrefix      string
	SuperAdminUser  string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type DatabaseSettingSV2 map[string]DatabaseSettingS

type TokenSettingS struct {
	Issuer  string
	Subject string
	Expires int
	Key     string
}
