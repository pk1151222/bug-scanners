package main

import (
	"fmt"
	"log"
	"os"
	"bug-scanner/subdomain"
	"bug-scanner/tlsanalysis"
	"bug-scanner/output"
	"bug-scanner/pdfreport"
)

func main() {
	// Example domain input
	domain := "example.com"

	// Step 1: Subdomain Enumeration
	subdomains, err := subdomain.EnumerateSubdomains(domain)
	if err != nil {
		log.Fatalf("Error enumerating subdomains: %v", err)
	}
	fmt.Println("Subdomains found:", subdomains)

	// Step 2: TLS Analysis
	tlsDetails, err := tlsanalysis.AnalyzeTLS(domain)
	if err != nil {
		log.Fatalf("Error analyzing TLS: %v", err)
	}
	fmt.Println("TLS Details:", tlsDetails)

	// Step 3: Save results to output file (JSON/CSV)
	err = output.SaveResults(subdomains, tlsDetails)
	if err != nil {
		log.Fatalf("Error saving results: %v", err)
	}
	fmt.Println("Results saved successfully!")

	// Step 4: Generate PDF report
	err = pdfreport.GeneratePDF(subdomains, tlsDetails)
	if err != nil {
		log.Fatalf("Error generating PDF report: %v", err)
	}
	fmt.Println("PDF report generated successfully!")
}
