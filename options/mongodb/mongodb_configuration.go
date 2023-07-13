package mongodb

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "mongodb"
)

type MongodbConfiguration struct {
	//列表
	MongodbList map[string]MongodbOptions `mapstructure:"list" json:"list" yaml:"list"`
}

func (c *MongodbConfiguration) GetDefaultOptions() *MongodbOptions {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *MongodbConfiguration) GetOptions(aliasName string) *MongodbOptions {
	if len(c.MongodbList) <= 0 {
		return nil
	}
	if len(aliasName) <= 0 {
		aliasName = AliasName_Default
	}
	item, ok := c.MongodbList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// 序列化为json字符串
func (c *MongodbConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (MongodbConfiguration, error) {
	var mongodbConfiguration MongodbConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &mongodbConfiguration)
	return mongodbConfiguration, err
}
