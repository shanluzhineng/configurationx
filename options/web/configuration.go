package web

import "encoding/json"

const (
	ConfigurationKey string = "web"
)

type Configuration struct {
	JWT *JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	// cors
	Cors       CORS                   `mapstructure:"cors" json:"cors" yaml:"cors"`
	Properties map[string]interface{} `mapstructure:"properties" json:"properties" yaml:"properties"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Cors: CORS{
			Mode:      CorsMode_AllowAll,
			Whitelist: make([]CORSWhitelist, 0),
		},
		Properties: make(map[string]interface{}),
	}
}

// serialize Configuration to json
func (c *Configuration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}
