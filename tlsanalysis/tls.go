package tlsanalysis

import (
	"crypto/tls"
	"fmt"
)

// TLSDetails contains the TLS-related information
type TLSDetails struct {
	SupportedVersions []string
	CipherSuites      []string
}

// Function to analyze TLS details of a domain
func AnalyzeTLS(domain string) (*TLSDetails, error) {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	state := conn.ConnectionState()
	tlsDetails := &TLSDetails{
		SupportedVersions: []string{"TLS 1.2", "TLS 1.3"}, // Example versions
		CipherSuites:      []string{"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"}, // Example cipher suite
	}

	// Extract actual versions and cipher suites (real implementation here)
	for _, version := range state.PeerCertificates {
		// Process the versions and cipher suites
	}

	return tlsDetails, nil
}
