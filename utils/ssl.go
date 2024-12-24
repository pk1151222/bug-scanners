// internal/utils/ssl.go

package utils

import (
	"crypto/sha256"
	"crypto/tls"
	"fmt"
)

// SSL प्रमाणपत्र का फिंगरप्रिंट प्राप्त करने का फ़ंक्शन।
func GetSSLCertificateFingerprint(domain string) (string, error) {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		return "", fmt.Errorf("डोमेन से कनेक्ट करने में त्रुटि: %v", err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	fingerprint := sha256.Sum256(cert.Raw)
	return fmt.Sprintf("%x", fingerprint), nil
}
