package consul

import "os"

var _options ConsulOptions

const (
	ConfigurationKey string = "consul"
)

// 获取consul配置
func GetConsulOptions() *ConsulOptions {
	return &_options
}

// // 从中读取配置
// func ReadFrom(v *viper.Viper) error {
// 	err := v.UnmarshalKey(ConfigurationKey, &_options)
// 	if err != nil {
// 		return err
// 	}
// 	(&_options).Normalize()
// 	if len(_options.AclToken) > 0 {
// 		os.Setenv("CONSUL_HTTP_TOKEN", _options.AclToken)
// 	}

// 	return nil
// }

func SetConsul(options *ConsulOptions) {
	_options = *options
	(&_options).Normalize()
	if len(_options.AclToken) > 0 {
		os.Setenv("CONSUL_HTTP_TOKEN", _options.AclToken)
	}
}
