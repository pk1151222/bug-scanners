package sql_scanner

import (
	"fmt"
	"net/http"
	"strings"
)

// TestPayloads for detecting SQL Injection
var TestPayloads = []string{
	"' OR 1=1 --",
	"' UNION SELECT NULL --",
	"\" OR \"1\"=\"1",
}

// ScanForSQLInjection checks for SQL injection vulnerabilities.
func ScanForSQLInjection(url string) ([]string, error) {
	var vulnerabilities []string

	for _, payload := range TestPayloads {
		targetURL := fmt.Sprintf("%s?test=%s", url, payload)
		resp, err := http.Get(targetURL)
		if err != nil {
			return nil, fmt.Errorf("error testing SQL Injection: %v", err)
		}
		defer resp.Body.Close()

		// Basic detection by looking for errors in the response
		if strings.Contains(resp.Status, "500") || strings.Contains(resp.Status, "error") {
			vulnerabilities = append(vulnerabilities, fmt.Sprintf("Potential SQL Injection with payload: %s", payload))
		}
	}

	return vulnerabilities, nil
}
