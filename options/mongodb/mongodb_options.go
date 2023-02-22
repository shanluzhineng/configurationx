package mongodb

import "time"

type MongodbOptions struct {
	Uri string `json:"uri,omitempty"`

	ConnectTimeout       *time.Duration `json:"connectTimeout,omitempty"`
	ConnectTimeoutSet    *bool          `json:"connectTimeoutSet,omitempty"`
	HeartbeatInterval    *time.Duration `json:"heartbeatInterval,omitempty"`
	HeartbeatIntervalSet *bool          `json:"heartbeatIntervalSet,omitempty"`
	Hosts                *[]string      `json:"hosts,omitempty"`
	MaxConnIdleTime      *time.Duration `json:"maxConnIdleTime,omitempty"`
	MaxConnIdleTimeSet   *bool          `json:"maxConnIdleTimeSet,omitempty"`
	MaxPoolSize          *uint64        `json:"maxPoolSize,omitempty"`
	MaxPoolSizeSet       *bool          `json:"maxPoolSizeSet,omitempty"`
	MinPoolSize          *uint64        `json:"minPoolSize,omitempty"`
	MinPoolSizeSet       *bool          `json:"minPoolSizeSet,omitempty"`
	MaxConnecting        *uint64        `json:"maxConnecting,omitempty"`
	MaxConnectingSet     *bool          `json:"maxConnectingSet,omitempty"`
	Password             *string        `json:"password,omitempty"`
	PasswordSet          *bool          `json:"passwordSet,omitempty"`

	//是否开启mongodb命令监控，默认值为false
	EnableCommandMonitor *bool `json:"enableCommandMonitor,omitempty"`
}
