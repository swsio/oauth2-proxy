package options

// AlphaOptions contains alpha structured configuration options.
// Usage of these options allows users to access alpha features that are not
// available as part of the primary configuration structure for OAuth2 Proxy.
//
// :::warning
// The options within this structure are considered alpha.
// They may change between releases without notice.
// :::
type AlphaOptions struct {
	// Upstreams is used to configure upstream servers.
	// Once a user is authenticated, requests to the server will be proxied to
	// these upstream servers based on the path mappings defined in this list.
	Upstreams Upstreams `json:"upstreams,omitempty"`

	// InjectRequestHeaders is used to configure headers that should be added
	// to request to upstream servers.
	// Headers may source values from either the authenticated user's session
	// or from a static secret value.
	InjectRequestHeaders []Header `json:"injectRequestHeaders,omitempty"`

	// InjectResponseHeaders is used to configure headers that should be added
	// to responses from the proxy.
	// This is typically used when using the proxy as an external authentication
	// provider in conjunction with another proxy such as NGINX and its
	// auth_request module.
	// Headers may source values from either the authenticated user's session
	// or from a static secret value.
	InjectResponseHeaders []Header `json:"injectResponseHeaders,omitempty"`
}
