package db

import "strings"

const (
	DbType_Mysql      string = "mysql"
	DbType_Sqlite     string = "sqlite"
	DbType_Sqlserver  string = "sqlserver"
	DbType_Postgresql string = "postgresql"
)

type GeneralDB struct {
	// 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	//主机名
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	//:端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 高级配置
	Config string `mapstructure:"config" json:"config" yaml:"config"`
	// 数据库名
	Dbname string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	// 数据库用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 数据库密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// 空闲中的最大连接数
	MaxIdleConns int `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	// 打开到数据库的最大连接数
	MaxOpenConns int `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	// 是否开启Gorm全局日志
	LogMode string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	// 是否通过zap写入日志文件
	LogZap bool `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
}

type SpecializedDB struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`

	Disable bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}

type DsnProvider interface {
	Dsn() string
}

func (m *GeneralDB) Dsn() string {
	if strings.EqualFold(m.DbType, DbType_Mysql) {
		return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
	}
	return ""
}

func (m *GeneralDB) GetLogMode() string {
	return m.LogMode
}
