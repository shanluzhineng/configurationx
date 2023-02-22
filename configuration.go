package configurationx

import (
	"errors"

	"github.com/abmpio/configurationx/options"
	"github.com/spf13/viper"
)

var (
	_instance *Configuration
)

type Configuration struct {
	Logger        viper.Logger
	viper         *viper.Viper
	ConfigManager ConfigManager

	BaseConsulPathList []string `json:"-"`
	EtcSubPath         string   `json:"-"`
	options.Options
}

type ConfigurationReadOption func(c *Configuration)

// 注册额外的配置信息
func RegistExtraProperties(key string, value interface{}) {
	options.RegistExtraProperties(key, value)
}

// 获取配置值
func GetOption(key string) interface{} {
	return GetInstance().Options.GetExtraProperties(key)
}

// 构建配置信息,将设置GetInstance()方法的返回值
func Use(c *Configuration, opts ...ConfigurationReadOption) *Configuration {
	for _, eachOpt := range opts {
		eachOpt(c)
	}

	_instance = c
	return _instance
}

// 获取实例
func GetInstance() *Configuration {
	return _instance
}

// Load出一个Configuration对象
func Load(etcSubPath string, opts ...ConfigurationReadOption) *Configuration {
	configuration := New()
	// configuration.Logger.Debug("准备初始化应用配置信息...")
	configuration.EtcSubPath = etcSubPath

	for _, eachOpt := range opts {
		if eachOpt == nil {
			continue
		}
		eachOpt(configuration)
	}

	//读取数据
	configuration.ReadFrom(configuration.viper)
	return configuration
}

func New() *Configuration {
	return NewConfiguration(viper.New())
}

func NewConfiguration(viper *viper.Viper) *Configuration {
	if viper == nil {
		panic(errors.New("viper参数不能为nil"))
	}
	configuration := new(Configuration)
	configuration.Logger = newMyLogger()
	configuration.viper = viper
	configuration.BaseConsulPathList = []string{"abmp/"}
	configuration.Options = options.NewOptions()
	return configuration
}

// Reset all configuration
func (c *Configuration) Reset() {
	c.viper = viper.New()
}

// Merge others to c
func (c *Configuration) Merge(opts ...ConfigurationReadOption) *Configuration {
	for _, eachOpt := range opts {
		if eachOpt == nil {
			continue
		}
		eachOpt(c)
	}
	return c
}

// 反序列化指定的key的值到一个对象中
func (c *Configuration) UnmarshFromKey(key string, v interface{}) error {
	return c.viper.UnmarshalKey(key, v)
}

func (c *Configuration) GetViper() *viper.Viper {
	return c.viper
}

// Merge from other Configuration
func (c *Configuration) MergeFrom(source *Configuration) *Configuration {
	c.viper.MergeConfigMap(source.viper.AllSettings())
	return c
}

func ReadFromConfiguration(source *Configuration) ConfigurationReadOption {
	return func(c *Configuration) {
		c.MergeFrom(source)
	}
}
