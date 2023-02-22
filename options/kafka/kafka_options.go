package kafka

import (
	"errors"
	"os"
	"time"
)

var (
	Error_Empty     error = errors.New("empty producer configuration")
	DefaultMinBytes int   = 1
	DefaultMaxBytes int   = 1024 * 1024

	ENV_AppName = "app.name"
)

// kafka配置
type KafkaOptions struct {
	DialOptions `mapstructure:",squash"`
	//key为topic,value为相关值
	ProducerList map[string]*ProducerOptions `mapstructure:"producers" json:"producers"`
	//key为topic,value为相关值
	ConsumerList map[string]*ConsumerOptions `mapstructure:"consumers" json:"consumers"`
}

// 获取指定key的生产者配置
func (o *KafkaOptions) GetProducer(key string) *ProducerOptions {
	if len(o.ProducerList) <= 0 {
		return nil
	}
	value, ok := o.ProducerList[key]
	if !ok {
		return nil
	}
	value.EnsureDefaultValueIfEmpty()
	return value
}

// 获取指定key的消息者配置
func (o *KafkaOptions) GetConsumer(key string) *ConsumerOptions {
	if len(o.ConsumerList) <= 0 {
		return nil
	}
	value, ok := o.ConsumerList[key]
	if !ok {
		value = &ConsumerOptions{}
	}
	value.EnsureDefaultValueIfEmpty()
	return value
}

type DialOptions struct {
	//使用的网络，正常情况下都是tcp
	Network string `mapstructure:"network" json:"network,omitempty"`
	// The list of brokers used to discover the partitions available on the
	// kafka cluster.
	Brokers []string `mapstructure:"brokers" json:"brokers"`
}

type ProducerOptions struct {
	ChunkSize     int           `json:"chunkSize,omitempty"`
	FlushInterval time.Duration `json:"flushInterval,omitempty"`
}

func (p *ProducerOptions) EnsureDefaultValueIfEmpty() {
	if p.ChunkSize <= 0 {
		p.ChunkSize = DefaultMaxBytes
	}
	if p.FlushInterval <= 0 {
		p.FlushInterval = time.Second
	}
}

type ConsumerOptions struct {
	Disabled bool   `json:"disabled,omitempty" mapstructure:"disabled,omitempty"`
	Group    string `json:"group,omitempty" mapstructure:"group,omitempty"`
	//消费者个数
	Consumers  int `json:"consumers,omitempty" mapstructure:"consumers,omitempty"`
	Processors int `json:"processors,omitempty" mapstructure:"processors,omitempty"`
	//最小字节，默认值为1,
	MinBytes int `json:"minBytes,omitempty" mapstructure:"minBytes,omitempty"`
	//最大字节数，默认值为10485760,10M
	MaxBytes int `json:"maxBytes,omitempty" mapstructure:"maxBytes,omitempty"`
}

func (c *ConsumerOptions) EnsureDefaultValueIfEmpty() {
	if len(c.Group) <= 0 {
		groupName := os.Getenv(ENV_AppName)
		if len(groupName) <= 0 {
			groupName, _ = os.Hostname()
		}
		c.Group = groupName
	}
	if c.Consumers <= 0 {
		c.Consumers = 1
	}
	if c.Processors <= 0 {
		c.Processors = 1
	}
	if c.MinBytes <= 0 {
		c.MinBytes = DefaultMinBytes
	}
	if c.MaxBytes <= c.MinBytes {
		c.MaxBytes = DefaultMaxBytes
	}
}
