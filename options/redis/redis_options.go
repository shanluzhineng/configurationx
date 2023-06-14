package redis

import "time"

type RedisOptions struct {
	KeyPrefix string `json:"keyPrefix,omitempty"`
	// The network type, either tcp or unix.
	// Default is tcp.
	Network string `json:"network,omitempty"`
	// host:port address.
	Addr string `json:"addr,omitempty"`

	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `json:"username,omitempty"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `json:"password,omitempty"`

	// Database to be selected after connecting to the server.
	DB int `json:"db,omitempty"`

	// Maximum number of retries before giving up.
	// Default is 3 retries; -1 (not 0) disables retries.
	MaxRetries *int `json:"maxRetries,omitempty"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff time.Duration `json:"minRetryBackoff,omitempty"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff time.Duration `json:"maxRetryBackoff,omitempty"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout *time.Duration `json:"dialTimeout,omitempty"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout *time.Duration `json:"readTimeout,omitempty"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout *time.Duration `json:"writeTimeout,omitempty"`

	// Type of connection pool.
	// true for FIFO pool, false for LIFO pool.
	// Note that fifo has higher overhead compared to lifo.
	PoolFIFO *bool `json:"poolFIFO,omitempty"`
	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize *int `json:"poolSize,omitempty"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns *int `json:"minIdleConns,omitempty"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge *time.Duration `json:"maxConnAge,omitempty"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout *time.Duration `json:"poolTimeout,omitempty"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout *time.Duration `json:"idleTimeout,omitempty"`
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency *time.Duration `json:"idleCheckFrequency,omitempty"`
}
