package casdoor

import (
	"io"
	"os"
	"strings"

	"github.com/spf13/viper"
)

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

	Disabled bool       `json:"disabled,omitempty"`
	Jwt      JwtOptions `json:"jwt,omitempty" mapstructure:"jwt"`
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

func (o *CasdoorOptions) Normalize() {
	if len(o.Certificate) <= 0 && strings.TrimSpace(o.CertificateFilePath) != "" {
		certData, err := readFile(o.CertificateFilePath)
		if err != nil {
			panic(err)
		}
		o.Certificate = string(certData)
	}
	o.Jwt.Normalize()
}

// read data from file
func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	return byteValue, err
}
