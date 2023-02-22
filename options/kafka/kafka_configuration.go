package kafka

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "kafka"
)

type KafkaConfiguration struct {
	//列表
	KafkaList map[string]KafkaOptions `mapstructure:"list" json:"list" yaml:"list"`
}

func (c *KafkaConfiguration) GetDefaultOptions() *KafkaOptions {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *KafkaConfiguration) GetOptions(aliasName string) *KafkaOptions {
	if len(c.KafkaList) <= 0 {
		return nil
	}
	item, ok := c.KafkaList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// 序列化为json字符串
func (c *KafkaConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (KafkaConfiguration, error) {
	var kafkaConfiguration KafkaConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &kafkaConfiguration)
	return kafkaConfiguration, err
}
