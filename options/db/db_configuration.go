package db

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "db"
)

type DbConfiguration struct {
	//列表
	DbList map[string]SpecializedDB `mapstructure:"list" json:"list" yaml:"list"`
}

// 获取主要的db配置
func (c *DbConfiguration) GetDefaultOptions() *SpecializedDB {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *DbConfiguration) GetOptions(aliasName string) *SpecializedDB {
	if len(c.DbList) <= 0 {
		return nil
	}
	item, ok := c.DbList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// 序列化为json字符串
func (c *DbConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (DbConfiguration, error) {
	var dbConfiguration DbConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &dbConfiguration)
	return dbConfiguration, err
}
