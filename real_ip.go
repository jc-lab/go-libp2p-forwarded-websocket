package websocket

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// GetRealIPFromHeader extracts the client's real IP address from HTTP request headers.
// It checks various proxy headers to find the actual IP.
func GetRealIPFromHeader(h http.Header) net.IP {
	// Check X-Real-IP header (used by Nginx and others)
	ipStr := h.Get("X-Real-IP")
	if ip := validateIp(ipStr); ip != nil {
		return ip
	}

	// Check X-Forwarded-For header (used by most proxies)
	// Format: client, proxy1, proxy2, ...
	ipStr = h.Get("X-Forwarded-For")
	if ipStr != "" {
		// Extract the first IP from the comma-separated list
		ips := strings.Split(ipStr, ",")
		for _, ipItem := range ips {
			ipItem = strings.TrimSpace(ipItem)
			if ip := validateIp(ipItem); ip != nil {
				return ip
			}
		}
	}

	// Check CF-Connecting-IP header (used by Cloudflare)
	ipStr = h.Get("CF-Connecting-IP")
	if ip := validateIp(ipStr); ip != nil {
		return ip
	}

	// Check True-Client-IP header (used by Akamai, Cloudflare, etc.)
	ipStr = h.Get("True-Client-IP")
	if ip := validateIp(ipStr); ip != nil {
		return ip
	}

	return nil
}

// validateIp checks if a string is a valid IP address
func validateIp(ip string) net.IP {
	if ip == "" {
		return nil
	}
	return net.ParseIP(ip)
}

func IpPort(ip net.IP, port int) string {
	if ip.To4() == nil {
		return fmt.Sprintf("[%s]:%d", ip.String(), port)
	}
	return fmt.Sprintf("%s:%d", ip.String(), port)
}
