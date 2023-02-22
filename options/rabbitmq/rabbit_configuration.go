package rabbitmq

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "rabbitmq"
)

type RabbitmqConfiguration struct {
	//RabbitmqOptions list ,key is aliasName
	RabbitmqList map[string]RabbitmqOptions `mapstructure:"list" json:"list" yaml:"list"`
}

func (c *RabbitmqConfiguration) GetDefaultOptions() *RabbitmqOptions {

	result := c.GetOptions("")
	if result == nil {
		result = c.GetOptions(AliasName_Default)
	}
	return result
}

// get RabbitmqOptions by aliasName
func (c *RabbitmqConfiguration) GetOptions(aliasName string) *RabbitmqOptions {
	if len(c.RabbitmqList) <= 0 {
		return nil
	}
	item, ok := c.RabbitmqList[aliasName]
	if !ok {
		return nil
	}
	return &item
}

// serialize RabbitmqConfiguration to json
func (c *RabbitmqConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// read RabbitmqConfiguration from viper instance
func ReadFrom(v *viper.Viper) (*RabbitmqConfiguration, error) {
	var rabbitmqConfiguration *RabbitmqConfiguration
	err := v.UnmarshalKey(ConfigurationKey, rabbitmqConfiguration)
	return rabbitmqConfiguration, err
}
