package redis

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "redis"
)

type RedisConfiguration struct {
	//列表
	RedisList map[string]RedisOptions `mapstructure:"list" json:"list" yaml:"list"`
}

func (c *RedisConfiguration) GetDefaultOptions() *RedisOptions {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *RedisConfiguration) GetOptions(aliasName string) *RedisOptions {
	if len(c.RedisList) <= 0 {
		return nil
	}
	item, ok := c.RedisList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// 序列化为json字符串
func (c *RedisConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (RedisConfiguration, error) {
	var redisConfiguration RedisConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &redisConfiguration)
	return redisConfiguration, err
}
