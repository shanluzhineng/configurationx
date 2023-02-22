package minio

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "minio"
)

type MinioConfiguration struct {
	//列表
	MinioList map[string]MinioOptions `mapstructure:"list" json:"list" yaml:"list"`
}

func (c *MinioConfiguration) GetDefaultOptions() *MinioOptions {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *MinioConfiguration) GetOptions(aliasName string) *MinioOptions {
	if len(c.MinioList) <= 0 {
		return nil
	}
	item, ok := c.MinioList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// 序列化为json字符串
func (c *MinioConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (MinioConfiguration, error) {
	var minioConfiguration MinioConfiguration

	err := v.UnmarshalKey(ConfigurationKey, &minioConfiguration)
	return minioConfiguration, err
}
