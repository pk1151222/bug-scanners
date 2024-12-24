package utils

import (
	"crypto/tls"
	"fmt"
	"net"
)

func CheckTLS(domain string) ([]string, error) {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	tlsInfo := conn.ConnectionState()
	var tlsVersions []string
	for _, version := range tlsInfo.PeerCertificates {
		tlsVersions = append(tlsVersions, fmt.Sprintf("%v", version.SignatureAlgorithm))
	}
	return tlsVersions, nil
}
