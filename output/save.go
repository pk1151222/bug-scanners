package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveResults saves the subdomain and TLS details in JSON format
func SaveResults(subdomains []string, tlsDetails *TLSDetails) error {
	file, err := os.Create("results.json")
	if err != nil {
		return err
	}
	defer file.Close()

	results := map[string]interface{}{
		"subdomains":    subdomains,
		"tls_details": tlsDetails,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		return err
	}
	return nil
}
