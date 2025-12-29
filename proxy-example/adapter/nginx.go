package adapter

import (
	"fmt"
	"proxy-example/domain"
)

// Nginx is a proxy with basic rate limiting.
type Nginx struct {
	application        domain.Server
	maxAllowedRequests int
	rateLimiter        map[string]int
	logger             domain.Logger
}

// NewNginx builds an Nginx proxy with defaults.
func NewNginx(app domain.Server, logger domain.Logger) *Nginx {
	return &Nginx{
		application:        app,
		maxAllowedRequests: 2,
		rateLimiter:        make(map[string]int),
		logger:             logger,
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		n.logger.Log(fmt.Sprintf("[Nginx] Blocked request to %s (Rate Limit Exceeded)", url))
		return 403, "Forbidden"
	}
	n.logger.Log(fmt.Sprintf("[Nginx] Forwarding request to %s", url))
	return n.application.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] >= n.maxAllowedRequests {
		return false
	}
	n.rateLimiter[url]++
	return true
}
