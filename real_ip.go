package websocket

import (
	"net"
	"net/http"
	"strings"
)

// GetRealIPFromHeader extracts the client's real IP address from HTTP request headers.
// It checks various proxy headers to find the actual IP.
func GetRealIPFromHeader(h http.Header) string {
	// Check X-Real-IP header (used by Nginx and others)
	ip := h.Get("X-Real-IP")
	if ip != "" && isValidIP(ip) {
		return ip
	}

	// Check X-Forwarded-For header (used by most proxies)
	// Format: client, proxy1, proxy2, ...
	ip = h.Get("X-Forwarded-For")
	if ip != "" {
		// Extract the first IP from the comma-separated list
		ips := strings.Split(ip, ",")
		for _, ipItem := range ips {
			ipItem = strings.TrimSpace(ipItem)
			if ipItem != "" && isValidIP(ipItem) {
				return ipItem
			}
		}
	}

	// Check CF-Connecting-IP header (used by Cloudflare)
	ip = h.Get("CF-Connecting-IP")
	if ip != "" && isValidIP(ip) {
		return ip
	}

	// Check True-Client-IP header (used by Akamai, Cloudflare, etc.)
	ip = h.Get("True-Client-IP")
	if ip != "" && isValidIP(ip) {
		return ip
	}

	return ""
}

// isValidIP checks if a string is a valid IP address
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
