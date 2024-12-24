// internal/scanner/scanner.go

package scanner

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
	"github.com/pk1151222/bug-scanners/internal/models"
)

// डोमेन के लिए SNI बग स्कैन करने का फ़ंक्शन।
func ScanSNI(domain string) (*models.ScanResult, error) {
	result := &models.ScanResult{Domain: domain}

	// IP पता रिजॉल्व करें
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, fmt.Errorf("डोमेन %s के लिए IP रिजॉल्व करने में त्रुटि: %v", domain, err)
	}
	result.IP = ips[0].String()

	// TLS हैंडशेक करने के लिए SNI बग जांचें
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{ServerName: domain})
	if err != nil {
		return nil, fmt.Errorf("डोमेन %s से कनेक्ट होने में त्रुटि: %v", domain, err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	result.Server = state.ServerName

	// अन्य TLS विवरण (Cipher Suites, Versions) जोड़ें
	// result.TLSVersions, result.CipherSuites, आदि भरें।

	return result, nil
}
