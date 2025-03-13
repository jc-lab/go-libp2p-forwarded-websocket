package websocket

import (
	"net"
	"testing"
)

func TestIpPort(t *testing.T) {
	// Test cases structure
	tests := []struct {
		name     string
		ip       string
		port     int
		expected string
	}{
		{
			name:     "IPv4 standard address",
			ip:       "192.168.1.1",
			port:     8080,
			expected: "192.168.1.1:8080",
		},
		{
			name:     "IPv4 localhost",
			ip:       "127.0.0.1",
			port:     80,
			expected: "127.0.0.1:80",
		},
		{
			name:     "IPv4 with zero port",
			ip:       "10.0.0.1",
			port:     0,
			expected: "10.0.0.1:0",
		},
		{
			name:     "IPv6 standard address",
			ip:       "2001:db8::1",
			port:     8080,
			expected: "[2001:db8::1]:8080",
		},
		{
			name:     "IPv6 localhost",
			ip:       "::1",
			port:     80,
			expected: "[::1]:80",
		},
		{
			name:     "IPv6 with zero port",
			ip:       "2001:db8::1",
			port:     0,
			expected: "[2001:db8::1]:0",
		},
		{
			name:     "IPv6 full format",
			ip:       "2001:0db8:0000:0000:0000:0000:0000:0001",
			port:     8080,
			expected: "[2001:db8::1]:8080",
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := net.ParseIP(tt.ip)
			if ip == nil {
				t.Fatalf("failed to parse IP address: %s", tt.ip)
			}

			result := IpPort(ip, tt.port)
			if result != tt.expected {
				t.Errorf("IpPort(%v, %d) = %s; want %s",
					tt.ip, tt.port, result, tt.expected)
			}
		})
	}
}
