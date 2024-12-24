// internal/models/scan_result.go

package models

// ScanResult डोमेन स्कैन के परिणामों को संग्रहित करता है।
type ScanResult struct {
	Domain       string
	IP           string
	Server       string
	Ports        []int
	TLSVersions  []string
	CipherSuites []string
	ALPNProtocols []string
	ESNI         bool
	ECH          bool
}
