package web

type CorsMode string

const (
	CorsMode_AllowAll  = "allow-all"
	CorsMode_Whitelist = "whitelist"
)

type CORS struct {
	Mode      CorsMode        `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowedOrigins string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
}

func (c *CORS) GetAllowedOrigins() []string {
	if len(c.Whitelist) <= 0 {
		return make([]string, 0)
	}
	list := make([]string, 0)
	for _, eachWhitelist := range c.Whitelist {
		list = append(list, eachWhitelist.AllowedOrigins)
	}
	return list
}
