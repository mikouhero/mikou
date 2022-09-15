package setting

import "time"

//  ServerSettingS configs/config.yaml  server 节点信息映射
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//  ServerSettingS configs/config.yaml  App 节点信息映射
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

//  DatabaseSettingS configs/config.yaml  Database 节点信息映射
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

//  DatabaseSettingSV2 configs/config.yaml  Databasev2 节点信息映射
type DatabaseSettingSV2 map[string]DatabaseSettingS

//  TokenSettingS configs/config.yaml  token 节点信息映射
type TokenSettingS struct {
	Issuer  string
	Subject string
	Expires int
	Key     string
}
