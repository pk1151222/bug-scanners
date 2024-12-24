package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"bug-scanners/internal/subdomain"
	"bug-scanners/internal/scanner"
	"bug-scanners/internal/utils"
)

func main() {
	// Example domain or list of domains to scan
	domains := []string{"example.com", "anotherexample.com"}

	// Create a results file to save scan results
	resultsFile, err := os.Create("scan_results.json")
	if err != nil {
		log.Fatalf("Error creating result file: %v", err)
	}
	defer resultsFile.Close()

	var results []scanner.ScanResult

	// Iterate through each domain
	for _, domain := range domains {
		// Enumerate subdomains
		subdomains, err := subdomain.EnumerateSubdomains(domain)
		if err != nil {
			log.Printf("Error enumerating subdomains for %s: %v", domain, err)
			continue
		}

		// Scan TLS/SSL and other vulnerabilities for each subdomain
		for _, subdomain := range subdomains {
			tlsVersions, err := scanner.ScanTLS(subdomain)
			if err != nil {
				log.Printf("Error scanning TLS for %s: %v", subdomain, err)
				continue
			}

			// Save the scan result
			result := scanner.ScanResult{
				Domain:      subdomain,
				TLSVersions: tlsVersions,
				Timestamp:   time.Now(),
			}
			results = append(results, result)
		}
	}

	// Save scan results to JSON file
	err = utils.SaveResults(resultsFile, results)
	if err != nil {
		log.Printf("Error saving results: %v", err)
	}

	fmt.Println("Scan completed, results saved.")
}
