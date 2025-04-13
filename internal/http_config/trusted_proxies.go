// File: internal/http_config/trusted_proxies.go
// Description: This package contains the configuration for trusted proxies in the Gin HTTP server.
// It sets the trusted proxies to allow the server to correctly handle forwarded headers.
// It is used to prevent IP spoofing and ensure that the server can trust the incoming requests.
// It is important to configure trusted proxies correctly to avoid security vulnerabilities.
package http_config

import (
	"log"

	"github.com/gin-gonic/gin"
)

// SetTrustedProxies configures the trusted proxies for the Gin HTTP server.
// It sets the trusted proxies to allow the server to correctly handle forwarded headers.
func SetTrustedProxies(r *gin.Engine) {
	trustedProxies := []string{"127.0.0.1", "::1", "192.168.0.0/16", "172.16.0.0/8"}

	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
}
