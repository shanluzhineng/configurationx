package casdoor

const (
	// DefaultContextKey jwt
	DefaultContextKey = "jwt"
)

type JwtOptions struct {
	// The name of the property in the request where the user (&token) information
	// from the JWT will be stored.
	// Default value: "jwt"
	ContextKey string
	// A boolean indicating if the credentials are required or not
	// Default value: false
	CredentialsOptional bool
	// When set, all requests with the OPTIONS method will use authentication
	// if you enable this option you should register your route with iris.Options(...) also
	// Default: false
	EnableAuthOnOptions bool
	// When set, the expiration time of token will be check every time
	// if the token was expired, expiration error will be returned
	// Default: false
	Expiration bool
}

func (o *JwtOptions) Normalize() {
	if o.ContextKey == "" {
		o.ContextKey = DefaultContextKey
	}
}
