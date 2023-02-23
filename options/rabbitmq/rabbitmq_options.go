package rabbitmq

import "time"

type RabbitmqOptions struct {
	DialOptions `mapstructure:",squash"`
}

// rabbitmq connection options
type DialOptions struct {
	//rabbitmq connection string, example:amqp://guest:guest@localhost:5672/
	RawUrl string `mapstructure:"rawUrl" json:"rawUrl"`

	// Vhost specifies the namespace of permissions, exchanges, queues and
	// bindings on the server.  Dial sets this to the path parsed from the URL.
	Vhost string `mapstructure:"vhost,omitempty" json:"vhost,omitempty"`
	//0 max channels means 2^16 - 1
	ChannelMax int `mapstructure:"channelMax,omitempty" json:"channelMax,omitempty"`
	//0 max bytes means unlimited
	FrameSize int `mapstructure:"frameSize,omitempty" json:"frameSize,omitempty"`
	// less than 1s uses the server's interval
	Heartbeat time.Duration `mapstructure:"heartbeat,omitempty" json:"heartbeat,omitempty"`
	// Properties is table of properties that the client advertises to the server.
	// This is an optional setting - if the application does not set this,
	// the underlying library will use a generic set of client properties.
	Properties map[string]interface{} `mapstructure:",remain"`

	// Connection locale that we expect to always be en_US
	// Even though servers must return it as per the AMQP 0-9-1 spec,
	// we are not aware of it being used other than to satisfy the spec requirements
	Locale string `mspstructure:"locale,omitempty" json:"locale,omitempty"`
}
