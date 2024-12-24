package scanner

import (
    "fmt"
    "github.com/pk1151222/bug-scanners/internal/subdomain"
    "github.com/pk1151222/bug-scanners/pkg/cve_scanner"
    "github.com/pk1151222/bug-scanners/pkg/security_headers"
    "github.com/pk1151222/bug-scanners/tlsanalysis"
)

func StartScan(args []string) {
    target := args[0] // Assuming first argument is the target URL

    // Scan for subdomains
    subdomains := subdomain.EnumerateSubdomains(target)
    fmt.Println("Subdomains found:", subdomains)

    // Scan for CVEs
    cves := cve_scanner.CheckCVE(target)
    fmt.Println("CVEs found:", cves)

    // Scan for security headers
    headers := security_headers.AnalyzeHeaders(target)
    fmt.Println("Security Headers:", headers)

    // TLS analysis
    tlsResults := tlsanalysis.CheckTLS(target)
    fmt.Println("TLS analysis:", tlsResults)
}
