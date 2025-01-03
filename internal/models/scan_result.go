// internal/models/scan_result.go

package models

// ScanResult डोमेन स्कैन के परिणामों को संग्रहित करता है।
type ScanResult struct {
	Domain       string
	IP           string
	Server       string
	Ports        []int
	TLSVersions  []string
    URL                string   `json:"url"`
    Subdomains         []string `json:"subdomains"`
    CVEs               []string `json:"cves"`
    XSSVulnerabilities []string
    WAFStatus string `json:"waf_status"`
`json:"xss_vulnerabilities"`
    SecurityHeaders    []string `json:"security_headers"`
SQLVulnerabilities []string `json:"sql_vulnerabilities"`

    // Add more vulnerability types here
	CipherSuites []string
	ALPNProtocols []string
	ESNI         bool
	ECH          bool
}
