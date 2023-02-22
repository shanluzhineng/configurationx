package minio

type MinioOptions struct {
	Endpoint         string `json:"endpoint"`
	MinioCredentials `mapstructure:"credentials"`

	Region   string `json:"region,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

type MinioCredentials struct {
	// AWS Access key ID
	AccessKeyID string `json:"accessKeyID,omitempty"`
	// AWS Secret Access Key
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
	// AWS Session Token
	SessionToken string `json:"sessionToken,omitempty"`
	// Signature Type.
	SignerType *int `json:"signerType,omitempty"`
}
