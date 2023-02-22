package elasticsearch

import "time"

// es配置
type ElasticsearchOptions struct {
	//A list of Elasticsearch nodes to use
	Addresses []string `json:"addresses"`
	//Username for HTTP Basic Authentication.
	Username string `json:"username,omitempty"`
	//Password for HTTP Basic Authentication.
	Password string `json:"password,omitempty"`
	Debuglog bool   `json:"debuglog"`

	//Base64-encoded token for authorization; if set, overrides username/password and service token.
	APIKey *string `json:"apiKey,omitempty"`
	//Service token for authorization; if set, overrides username/password.
	ServiceToken           *string `json:"serviceToken,omitempty"`
	CertificateFingerprint *string `json:"certificateFingerprint,omitempty"`

	// List of status codes for retry. Default: 502, 503, 504.
	RetryOnStatus *[]int `json:"retryOnStatus,omitempty"`
	// Default: false.
	DisableRetry *bool `json:"disableRetry,omitempty"`
	// Default: 3.
	MaxRetries *int `json:"maxRetries,omitempty"`

	// Default: false.
	CompressRequestBody *bool `json:"compressRequestBody,omitempty"`

	// Discover nodes when initializing the client. Default: false.
	DiscoverNodesOnStart *bool `json:"discoverNodesOnStart,omitempty"`
	// Discover nodes periodically. Default: disabled.
	DiscoverNodesInterval *time.Duration `json:"discoverNodesInterval,omitempty"`

	// Enable the metrics collection.
	EnableMetrics *bool `json:"enableMetrics,omitempty"`
	// Enable the debug logging.
	EnableDebugLogger *bool `json:"enableDebugLogger,omitempty"`
	// Enable sends compatibility header
	EnableCompatibilityMode *bool `json:"enableCompatibilityMode,omitempty"`

	// Disable the additional "X-Elastic-Client-Meta" HTTP header.
	DisableMetaHeader *bool `json:"disableMetaHeader,omitempty"`

	//request timeout,unit: Millisecond
	RequestTimout *time.Duration `json:"requestTimout,omitempty"`
}
