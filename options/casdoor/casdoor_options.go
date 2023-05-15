package casdoor

import "github.com/spf13/viper"

const (
	ConfigurationKey string = "casdoor"
)

type CasdoorOptions struct {
	Endpoint         string `json:"endpoint,omitempty"`
	ClientId         string `json:"clientId,omitempty"`
	ClientSecret     string `json:"clientSecret,omitempty"`
	Certificate      string `json:"certificate,omitempty"`
	OrganizationName string `json:"organizationName,omitempty"`
	ApplicationName  string `json:"applicationName,omitempty"`

	Jwt JwtOptions `json:"jwt,omitempty" mapstructure:"jwt"`
	// file path for Certificate
	CertificateFilePath string `json:"certificateFilePath,omitempty"`
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (CasdoorOptions, error) {
	var options CasdoorOptions

	err := v.UnmarshalKey(ConfigurationKey, &options)
	if err == nil {
		options.Jwt.Normalize()
	}
	return options, err
}
