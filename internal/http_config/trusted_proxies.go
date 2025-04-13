package http_config

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetTrustedProxies(r *gin.Engine) {
	trustedProxies := []string{"127.0.0.1", "::1", "192.168.0.0/16", "172.16.0.0/8"}

	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
}
