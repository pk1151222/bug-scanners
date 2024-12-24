package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
"github.com/pk1151222/bug-scanners/internal/xss_scanner"
	"github.com/pk1151222/bug-scanners/cmd"
	"github.com/pk1151222/bug-scanners/internal/models"
	"github.com/pk1151222/bug-scanners/internal/scanner"
	"github.com/pk1151222/bug-scanners/internal/subdomain"
	"github.com/pk1151222/bug-scanners/internal/utils"
	"github.com/pk1151222/bug-scanners/pkg/cve_scanner"
	"github.com/pk1151222/bug-scanners/pkg/security_headers"
)

func main() {
	// Initialize a Gin router for a simple web UI
	router := gin.Default()

	// Root route: Basic info
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Bug Scanners API. Start scanning by calling /scan endpoint.",
		})
	})

	// Scan URL route
	router.GET("/scan/:url", func(c *gin.Context) {
		url := c.Param("url")

// XSS vulnerability scan
        xssVulnerabilities, err := xss_scanner.ScanForXSS(url)
        if err != nil {
            fmt.Println("Error during XSS scan:", err)
        }
        fmt.Println("XSS Vulnerabilities:", xssVulnerabilities)

        // Combine all scan

		// Subdomain enumeration
		fmt.Println("Starting subdomain enumeration for:", url)
		subdomains := subdomain.EnumerateSubdomains(url)
		fmt.Println("Subdomains found:", subdomains)

		// CVE scanning
		fmt.Println("Starting CVE scanning for:", url)
		cves := cve_scanner.ScanCVE(url)
		fmt.Println("CVEs detected:", cves)

		// Security headers analysis
		headers := security_headers.AnalyzeSecurityHeaders(url)
		fmt.Println("Security Headers:", headers)

		// Scan result model
		scanResult := models.ScanResult{
			URL:               url,
			Subdomains:        subdomains,
			CVEs:              cves,
			SecurityHeaders:   headers,
			HTTPVulnerabilities: []models.Vulnerability{},  // Example for HTTP vulnerabilities
			SSLVulnerabilities: []models.Vulnerability{},   // Example for SSL vulnerabilities
		}

		// Save or process scan result here (can be saved to a database, file, etc.)
		// For now, print the result
		fmt.Println("Scan Result:", scanResult)

		// Return the scan result as a JSON response
		c.JSON(200, gin.H{
			"scan_result": scanResult,
		})
	})

	// Start the web server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
