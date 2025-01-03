package xss_scanner

import (
    "fmt"
    "net/http"
    "strings"
)

// ScanForXSS checks for potential XSS vulnerabilities in the provided URL
func ScanForXSS(url string) ([]string, error) {
    var vulnerabilities []string
    
    // Make a simple GET request to the URL
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching URL: %v", err)
    }
    defer resp.Body.Close()

    // Check if the response contains common XSS vectors
    xssPatterns := []string{"<script>", "</script>", "alert(", "onerror=", "eval("}
    body := resp.Body

    // Simple string search for XSS patterns
    for _, pattern := range xssPatterns {
        if strings.Contains(body, pattern) {
            vulnerabilities = append(vulnerabilities, fmt.Sprintf("Potential XSS vulnerability found with pattern: %s", pattern))
        }
    }
    return vulnerabilities, nil
}
