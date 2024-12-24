// pkg/cve_scanner/cve_checker.go

package cve_scanner

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type CVE struct {
    ID          string  `json:"id"`
    Title       string  `json:"title"`
    Severity    string  `json:"severity"`
    Description string  `json:"description"`
    CVSSScore   float64 `json:"cvss_score"`
    Exploitable bool    `json:"exploitable"`
}

func FetchCVEData(url string) ([]CVE, error) {
    cves := []CVE{
        {"CVE-2021-12345", "SQL Injection", "High", "SQL Injection vulnerability found.", 9.8, true},
        {"CVE-2021-67890", "XSS", "Medium", "Cross-site scripting vulnerability.", 5.4, false},
    }
    return cves, nil
}

func PrioritizeCVEs(cves []CVE) []CVE {
    var highPriority []CVE
    for _, cve := range cves {
        if cve.CVSSScore >= 7.0 {
            highPriority = append(highPriority, cve)
        }
    }
    return highPriority
}

func ScanCVE(url string) {
    cves, err := FetchCVEData(url)
    if err != nil {
        fmt.Println("Error fetching CVEs:", err)
        return
    }

    highPriorityCVEs := PrioritizeCVEs(cves)

    fmt.Println("High Priority CVEs:")
    for _, cve := range highPriorityCVEs {
        fmt.Println("CVE ID:", cve.ID)
        fmt.Println("Title:", cve.Title)
        fmt.Println("Severity:", cve.Severity)
        fmt.Println("CVSS Score:", cve.CVSSScore)
        fmt.Println("Exploitable:", cve.Exploitable)
    }
}
