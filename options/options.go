package options

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/abmpio/configurationx/options/consul"
	"github.com/abmpio/configurationx/options/db"
	"github.com/abmpio/configurationx/options/elasticsearch"
	"github.com/abmpio/configurationx/options/kafka"
	"github.com/abmpio/configurationx/options/minio"
	"github.com/abmpio/configurationx/options/mongodb"
	"github.com/abmpio/configurationx/options/rabbitmq"
	"github.com/abmpio/configurationx/options/redis"
	"github.com/abmpio/configurationx/options/web"
	"github.com/spf13/viper"
)

type Options struct {
	// db配置
	Db            *db.DbConfiguration
	Mongodb       *mongodb.MongodbConfiguration
	Redis         *redis.RedisConfiguration
	Minio         *minio.MinioConfiguration
	Elasticsearch *elasticsearch.ElasticsearchConfiguration
	Kafka         *kafka.KafkaConfiguration
	Consul        *consul.ConsulOptions
	Rabbitmq      *rabbitmq.RabbitmqConfiguration
	Web           *web.Configuration

	//其它属性
	extraProperties map[string]interface{}
}

const (
	_viperKeyDelim = "."
)

var (
	_registExtraPropertes map[string]interface{} = make(map[string]interface{})
)

// 注册额外的配置信息
func RegistExtraProperties(key string, value interface{}) {
	if key == "" {
		return
	}
	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		//必须传入一个指针
		panic("value必须是一个指针")
	}
	if value == nil || reflect.ValueOf(value).IsNil() {
		//空值，则执行删除
		delete(_registExtraPropertes, key)
		return
	}

	_registExtraPropertes[strings.ToLower(key)] = value
}

// 创建一个新的Options对象
func NewOptions() Options {
	return Options{
		Db:              &db.DbConfiguration{},
		Mongodb:         &mongodb.MongodbConfiguration{},
		Redis:           &redis.RedisConfiguration{},
		Minio:           &minio.MinioConfiguration{},
		Elasticsearch:   &elasticsearch.ElasticsearchConfiguration{},
		Kafka:           &kafka.KafkaConfiguration{},
		Consul:          &consul.ConsulOptions{},
		Rabbitmq:        &rabbitmq.RabbitmqConfiguration{},
		Web:             web.NewConfiguration(),
		extraProperties: make(map[string]interface{}),
	}
}

// 获取额外的属性值,如果key不存在，则直接返回nil
// 本方法不是线程安全的
func (o *Options) GetExtraProperties(key string) interface{} {
	if len(key) <= 0 {
		return nil
	}
	key = strings.ToLower(key)
	return o.extraProperties[key]
}

// 设置额外属性值
// 本方法不是线程安全的
func (o *Options) SetExtraProperties(key string, v interface{}) {
	if len(key) <= 0 {
		return
	}
	key = strings.ToLower(key)
	if v == nil {
		delete(o.extraProperties, key)
	} else {
		o.extraProperties[key] = v
	}
}

// 将额外的属性的key设置到一个对象中
// v必须是一个指针对象
func (o *Options) UnmarshalPropertiesTo(key string, v interface{}) bool {
	if v == nil {
		return false
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return false
	}
	//需要转换为小写
	key = strings.ToLower(key)
	propertiesValue := o.GetExtraProperties(key)
	if propertiesValue == nil {
		//key不存在，则直接返回
		return false
	}
	jsonValue, err := json.Marshal(propertiesValue)
	if err != nil {
		return false
	}
	err = json.Unmarshal(jsonValue, v)
	return err == nil
}

// 从中读取配置
func (o *Options) ReadFrom(v *viper.Viper) (err error) {
	_registExtraPropertes = make(map[string]interface{})

	//先设置外部注册的默认的值
	for eachKey, eachValue := range _registExtraPropertes {
		o.SetExtraProperties(eachKey, eachValue)
	}
	rootKeys := getAllRootKeysFromViper(v)
	if len(rootKeys) <= 0 {
		return nil
	}
	knowedOptions := map[string]interface{}{
		db.ConfigurationKey:            o.Db,
		mongodb.ConfigurationKey:       o.Mongodb,
		redis.ConfigurationKey:         o.Redis,
		minio.ConfigurationKey:         o.Minio,
		elasticsearch.ConfigurationKey: o.Elasticsearch,
		kafka.ConfigurationKey:         o.Kafka,
		consul.ConfigurationKey:        o.Consul,
		rabbitmq.ConfigurationKey:      o.Rabbitmq,
		web.ConfigurationKey:           o.Web,
	}
	//读取已知的配置key
	for eachKey, eachValue := range knowedOptions {
		err = v.UnmarshalKey(eachKey, eachValue)
		if err != nil {
			return err
		}
	}
	//标准化consul配置，以补充相关默认值
	(o.Consul).Normalize()

	// read consul options
	consul.SetConsul(o.Consul)
	//读取额外的配置key
	allRootKey := getAllRootKeysFromViper(v)
	for _, eachKey := range allRootKey {
		_, ok := knowedOptions[eachKey]
		if ok {
			continue
		}
		var currentValue interface{} = o.GetExtraProperties(eachKey)
		err = v.UnmarshalKey(eachKey, &currentValue)
		if err != nil {
			continue
		}
		//没注册，则直接设置
		o.SetExtraProperties(eachKey, currentValue)
	}
	return nil
}

// serialize Options to json string
func (c *Options) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// print Options using json
func (c *Options) PrintJsonString() {
	if c.Db != nil {
		fmt.Printf("db:%s \r\n", c.Db.ToJsonString())
	}
	if c.Mongodb != nil {
		fmt.Printf("mongodb:%s \r\n", c.Mongodb.ToJsonString())
	}
	if c.Redis != nil {
		fmt.Printf("redis:%s \r\n", c.Redis.ToJsonString())
	}
	if c.Minio != nil {
		fmt.Printf("minio:%s \r\n", c.Minio.ToJsonString())
	}
	if c.Elasticsearch != nil {
		fmt.Printf("elasticsearch:%s \r\n", c.Elasticsearch.ToJsonString())
	}
	if c.Kafka != nil {
		fmt.Printf("kafka:%s \r\n", c.Kafka.ToJsonString())
	}
	if c.Rabbitmq != nil {
		fmt.Printf("rabbitmq:%s \r\n", c.Rabbitmq.ToJsonString())
	}
	if c.Web != nil {
		fmt.Printf("web:%s \r\n", c.Web.ToJsonString())
	}
}

func getAllRootKeysFromViper(v *viper.Viper) []string {
	rootKeys := make(map[string]string)
	for _, eachKey := range v.AllKeys() {
		keyPath := strings.Split(eachKey, _viperKeyDelim)
		if len(keyPath) <= 0 {
			continue
		}
		if len(keyPath[0]) <= 0 {
			continue
		}
		rootKeys[keyPath[0]] = keyPath[0]
	}
	allRootKeys := make([]string, 0)
	for eachKey := range rootKeys {
		allRootKeys = append(allRootKeys, eachKey)
	}
	return allRootKeys
}
