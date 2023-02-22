package elasticsearch

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "elasticsearch"
)

type ElasticsearchConfiguration struct {
	//列表
	ElasticsearchList map[string]ElasticsearchOptions `mapstructure:"list" json:"list" yaml:"list"`
}

// 获取主要的db配置
func (c *ElasticsearchConfiguration) GetDefaultOptions() *ElasticsearchOptions {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *ElasticsearchConfiguration) GetOptions(aliasName string) *ElasticsearchOptions {
	if len(c.ElasticsearchList) <= 0 {
		return nil
	}
	item, ok := c.ElasticsearchList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// 序列化为json字符串
func (c *ElasticsearchConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (ElasticsearchConfiguration, error) {
	var elasticsearchConfiguration ElasticsearchConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &elasticsearchConfiguration)
	return elasticsearchConfiguration, err
}
